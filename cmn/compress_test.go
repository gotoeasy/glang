package cmn

import (
	"testing"
)

func Test_gzip(t *testing.T) {
	Info(Gzip("f:\\222\\ddd.tar", "f:\\222\\ddd.tar.gz"))
	Info(UnGzip("f:\\222\\ddd.tar.gz", "f:\\222\\dddsss"))
}

func Test_snappy(t *testing.T) {
	src := `使用snappy算法压缩（压缩速度快，占用资源少，压缩比适当，重复多则压缩比大，适用于重复较多的文本压缩）`
	src = `DEBUG ==> Parameters: 1589491593233129481(String), XFD(String), B3特菜档口(String), 003B307620221(String), SC-B3-076(String), SC(String), 王二军(String), 82107(String), 2022-09-22(String), 50400.0000000000(BigDecimal), 6(Integer), 收入-租金(String), 603084(String), 2023-02-01(String), 2023-02-28(String), 4200.0000000000(BigDecimal), 4200.0000000000(BigDecimal), 4000.0000000000(BigDecimal), 200.0000000000(BigDecimal), 5(String), 2023(String), 2(String), 2023-02-01~2023-02-28(String)`
	src = `ERROR 处理TODO跟进人释放发生异常`

	srcBytes := StringToBytes(src)
	bt := Compress(srcBytes)
	Info("压缩前字节长", len(srcBytes), "，压缩前后字节长", len(bt))

	de := UnCompress(srcBytes)
	Info(BytesToString(de) == src)
}
