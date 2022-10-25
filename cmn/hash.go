package cmn

import "hash/crc32"

func HashCode(bytes []byte) uint32 {
	return crc32.ChecksumIEEE(bytes)
}
