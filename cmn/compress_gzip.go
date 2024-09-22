package cmn

import (
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

// 指定文件压缩为gzip文件（文件名不支持中文）
func Gzip(srcFile string, gzipFile string) error {
	inFile, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	MkdirAll(Dir(gzipFile))
	outFile, err := os.Create(gzipFile)
	if err != nil {
		return err
	}

	gz := gzip.NewWriter(outFile)
	defer gz.Close()

	gz.Name = FileName(srcFile)

	_, err = io.Copy(gz, inFile)
	return err
}

// 解压gzip文件到指定目录
func UnGzip(gzipPathFile string, destPath string) error {
	err := os.MkdirAll(destPath, 0666)
	if err != nil {
		return err
	}
	gzFile, err := os.Open(gzipPathFile)
	if err != nil {
		return err
	}
	defer gzFile.Close()

	gz, err := gzip.NewReader(gzFile)
	if err != nil {
		return err
	}
	defer gz.Close()

	outFile, err := os.Create(filepath.Join(destPath, gz.Header.Name))
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, gz)
	return err
}

// 用gzip压缩字节数组
func GzipBytes(bts []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)

	if _, err := gz.Write(bts); err != nil {
		return nil, err
	}

	if err := gz.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// 解压gzip字节数组
func UnGzipBytes(gzipBytes []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(gzipBytes))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	unGzipdBytes, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return unGzipdBytes, nil
}
