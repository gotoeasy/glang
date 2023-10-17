package cmn

import (
	"github.com/jxskiss/base62"
)

// Base62编码（同Base62Encode）
func Base62(bts []byte) string {
	return Base62Encode(bts)
}

// Base62编码
func Base62Encode(bts []byte) string {
	return base62.EncodeToString(bts)
}

// Base62解码
func Base62Decode(str string) ([]byte, error) {
	return base62.DecodeString(str)
}
