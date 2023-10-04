package cmn

import (
	"context"
	"io"
	"strings"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/client"
	"github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/relay"
	"github.com/multiformats/go-multiaddr"
)

type P2pRelayHost struct {
	Host host.Host
}

// 新建中继P2P节点
func NewP2pRelayHost(key string, port string) (*P2pRelayHost, error) {
	prvKey, _, err := crypto.GenerateEd25519Key(strings.NewReader(key + "-funcNewP2pRelayHostParameterKey"))
	if err != nil {
		return nil, err
	}
	multiaddrRelay, _ := multiaddr.NewMultiaddr("/ip4/0.0.0.0/tcp/" + port)
	relayHost, err := libp2p.New(
		libp2p.ListenAddrs(multiaddrRelay),
		libp2p.Identity(prvKey),
		libp2p.EnableRelay(),
	)
	if err != nil {
		return nil, err
	}

	_, err = relay.New(relayHost)
	if err != nil {
		return nil, err
	}

	return &P2pRelayHost{Host: relayHost}, nil
}

// 节点信息
func (p *P2pRelayHost) HostInfo() string {
	if p.Host == nil {
		return ""
	}
	hostId := p.Host.ID().Pretty()
	info := "节点ID=" + hostId
	addrs := p.Host.Addrs()
	for _, addr := range addrs {
		info += ("\r\n节点地址=" + addr.String())
	}
	return info
}

// 当前节点作为客户端连接到中继节点 /ip4/{ip}/tcp/{port}/p2p/{peerid}
func (p *P2pRelayHost) ConnectRelayHost(relayHostAddr string) error {
	serverAddr, err := multiaddr.NewMultiaddr(relayHostAddr)
	if err != nil {
		return err
	}
	serverInfo, err := peer.AddrInfoFromP2pAddr(serverAddr)
	if err != nil {
		return err
	}
	if err := p.Host.Connect(context.Background(), *serverInfo); err != nil { // 建立连接
		return err
	}
	_, err = client.Reserve(context.Background(), p.Host, *serverInfo)
	if err != nil {
		return err
	}
	return nil
}

// 设定处理器
func (p *P2pRelayHost) SetStreamHandler(uri string, handler network.StreamHandler) {
	p.Host.SetStreamHandler(protocol.ID(uri), handler)
}

// 向目标节点发起请求并返回响应结果，地址通常为 /ip4/{ip}/tcp/{port}/p2p/{peerid} 或 /p2p/{relayPeerid}/p2p-circuit/p2p/{peerid} 或 /p2p/{peerid}
func (p *P2pRelayHost) Request(targetHostAddr string, uri string, dataBytes []byte, timeout time.Duration) ([]byte, error) {
	// 使用WithTimeout创建一个有超时限制的context
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel() // 保证超时后释放资源

	// 连接到目标节点
	targetAddr, err := multiaddr.NewMultiaddr(targetHostAddr)
	if err != nil {
		return nil, err
	}
	targetAddrInfo, err := peer.AddrInfoFromP2pAddr(targetAddr)
	if err != nil {
		return nil, err
	}
	if err := p.Host.Connect(ctx, *targetAddrInfo); err != nil {
		return nil, err
	}

	// 新建一个临时的会话流
	stream, err := p.Host.NewStream(ctx, targetAddrInfo.ID, protocol.ID(uri))
	if err != nil {
		return nil, err
	}
	defer stream.Close()

	// 发送请求数据
	err = WriteBytesToStream(stream, dataBytes)
	if err != nil {
		return nil, err
	}

	// 接收请求数据
	return ReadBytesFromStream(stream)
}

// 写流
func WriteBytesToStream(stream network.Stream, dataBytes []byte) error {

	// 写长度
	requestLength := StringToUint32(IntToString(len(dataBytes)), 0)
	_, err := stream.Write(Uint32ToBytes(requestLength))
	if err != nil {
		return err
	}

	// 写内容
	_, err = stream.Write(dataBytes)
	if err != nil {
		return err
	}
	return nil
}

// 读流
func ReadBytesFromStream(stream network.Stream) ([]byte, error) {
	// 读长度
	prefix := make([]byte, 4)
	_, err := io.ReadFull(stream, prefix)
	if err != nil {
		return nil, err
	}
	messageLength := BytesToUint32(prefix)

	// 读内容
	message := make([]byte, messageLength)
	_, err = io.ReadFull(stream, message)
	if err != nil {
		return nil, err
	}
	return message, nil
}
