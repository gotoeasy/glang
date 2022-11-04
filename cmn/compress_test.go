package cmn

import (
	"testing"
)

func Test_gzip(t *testing.T) {
	Info(Gzip("f:\\222\\ddd.tar", "f:\\222\\ddd.tar.gz"))
	Info(UnGzip("f:\\222\\ddd.tar.gz", "f:\\222\\dddsss"))
}
