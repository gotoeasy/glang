package p2p

import (
	"context"
	"encoding/json"
	"io"
	"strings"

	"github.com/gotoeasy/glang/cmn"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/p2p/net/swarm"
	"github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/client"
	"github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/relay"
	"github.com/multiformats/go-multiaddr"
	"github.com/valyala/fasthttp"
)

type P2pRelayHost struct {
	Host host.Host
}

type FileDataModel struct {
	Success     bool
	Message     string
	ContentType string
	Data        []byte
	Gzip        bool
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
	hostId := p.Host.ID().String()
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
		p.Host.Network().(*swarm.Swarm).Backoff().Clear(serverInfo.ID) // 清除连接失败的缓存
		return err
	}
	_, err = client.Reserve(context.Background(), p.Host, *serverInfo)
	if err != nil {
		p.Host.Network().(*swarm.Swarm).Backoff().Clear(serverInfo.ID) // 清除连接失败的缓存
		return err
	}
	return nil
}

// 设定处理器
func (p *P2pRelayHost) SetStreamHandler(uri string, handler network.StreamHandler) {
	p.Host.SetStreamHandler(protocol.ID(uri), handler)
}

// 【注册】向目标节点发起请求并返回响应结果，地址通常为 /ip4/{ip}/tcp/{port}/p2p/{peerid} 或 /p2p/{relayPeerid}/p2p-circuit/p2p/{peerid} 或 /p2p/{peerid}
func (p *P2pRelayHost) Regist(targetHostAddr string, uri string, dataBytes []byte) ([]byte, error) {
	// 连接到目标节点
	targetAddr, err := multiaddr.NewMultiaddr(targetHostAddr)
	if err != nil {
		return nil, err
	}
	targetAddrInfo, err := peer.AddrInfoFromP2pAddr(targetAddr)
	if err != nil {
		return nil, err
	}
	if err := p.Host.Connect(context.Background(), *targetAddrInfo); err != nil {
		return nil, err
	}

	// 新建一个临时的会话流
	stream, err := p.Host.NewStream(network.WithUseTransient(context.Background(), "临时会话"), targetAddrInfo.ID, protocol.ID(uri))
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

// 【请求】向目标节点发起请求并返回响应结果，地址通常为 /ip4/{ip}/tcp/{port}/p2p/{peerid} 或 /p2p/{relayPeerid}/p2p-circuit/p2p/{peerid} 或 /p2p/{peerid}
func (p *P2pRelayHost) Request(c *fasthttp.RequestCtx, targetHostAddr string, uri string, dataBytes []byte) (int64, error) {
	// 连接到目标节点
	targetAddr, err := multiaddr.NewMultiaddr(targetHostAddr)
	if err != nil {
		return 0, err
	}
	targetAddrInfo, err := peer.AddrInfoFromP2pAddr(targetAddr)
	if err != nil {
		return 0, err
	}
	if err := p.Host.Connect(context.Background(), *targetAddrInfo); err != nil {
		return 0, err
	}

	// 新建一个临时的会话流
	stream, err := p.Host.NewStream(network.WithUseTransient(context.Background(), "临时会话"), targetAddrInfo.ID, protocol.ID(uri))
	if err != nil {
		return 0, err
	}
	defer stream.Close()

	// 发送请求数据
	err = WriteBytesToStream(stream, dataBytes)
	if err != nil {
		return 0, err
	}
	// 接收请求数据
	bts, err := ReadBytesFromStream(stream)
	if err != nil {
		return 0, err
	}

	fdm := &FileDataModel{}
	if err := json.Unmarshal(bts, fdm); err != nil {
		return 0, err
	}

	if fdm.Success && fdm.Data != nil {
		if c != nil {
			c.SetContentType(fdm.ContentType)
			c.SetBody(fdm.Data)
			c.SetStatusCode(200)
		}
		return int64(len(fdm.Data)), nil
	}

	// 读长度
	prefix := make([]byte, 4)
	_, err = io.ReadFull(stream, prefix)
	if err != nil {
		return 0, err
	}
	fileLength := int64(cmn.BytesToUint32(prefix))

	// 读内容
	if c != nil {
		writer := c.Response.BodyWriter()
		c.SetContentType(fdm.ContentType)
		_, err = io.CopyN(writer, stream, fileLength)
		if err != nil {
			return 0, err
		}
		c.SetStatusCode(200)
	}

	return fileLength, nil
}

// 当前节点作为客户端连接到指定地址节点 /ip4/{ip}/tcp/{port}/p2p/{peerid}
func ConnectHost(thisHost host.Host, relayHostAddr string) error {
	serverAddr, err := multiaddr.NewMultiaddr(relayHostAddr)
	if err != nil {
		return err
	}
	serverInfo, err := peer.AddrInfoFromP2pAddr(serverAddr)
	if err != nil {
		return err
	}
	if err := thisHost.Connect(context.Background(), *serverInfo); err != nil { // 建立连接
		return err
	}
	return nil
}

// 写流
func WriteBytesToStream(stream network.Stream, dataBytes []byte) error {

	// 写长度
	requestLength := cmn.StringToUint32(cmn.IntToString(len(dataBytes)), 0)
	_, err := stream.Write(cmn.Uint32ToBytes(requestLength))
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
	messageLength := cmn.BytesToUint32(prefix)

	// 读内容
	message := make([]byte, messageLength)
	_, err = io.ReadFull(stream, message)
	if err != nil {
		return nil, err
	}
	return message, nil
}
