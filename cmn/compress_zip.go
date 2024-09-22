package cmn

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// 压缩指定目录为指定的zip文件
func Zip(srcFileOrPath string, zipPathFile string) error {
	srcFileorpath := strings.ReplaceAll(srcFileOrPath, "\\", "/") // window得这么干！这是个神奇的地方，不信就试
	MkdirAll(Dir(zipPathFile))
	zipFile, err := os.Create(zipPathFile)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// lenPrefix := Len(filepath.Dir(srcFileorpath)) // 绝对路径除去末尾目录名后的长度
	lenPrefix := Len(srcFileorpath) // 绝对路径除去末尾目录名后的长度

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	return filepath.Walk(srcFileorpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if srcFileorpath == path {
			return nil
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = SubString(path, lenPrefix+1, Len(path))
		header.Extra = StringToBytes("file-size:" + strconv.FormatInt(info.Size(), 10) + ";custom:true ")

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}
		return nil
	})

}

// 解压指定zip文件
func UnZip(zipFile string, destPath string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		fName, err := decodeGBK(f.Name)
		if err != nil {
			fName = f.Name
		}

		path := filepath.Join(destPath, fName)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func UnZipBytes(bts []byte, destPath string) error {
	zipReader, err := zip.NewReader(bytes.NewReader(bts), int64(len(bts)))
	if err != nil {
		return err
	}

	for _, f := range zipReader.File {
		fName, err := decodeGBK(f.Name)
		if err != nil {
			fName = f.Name
		}

		path := filepath.Join(destPath, fName)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
				return err
			}

			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, rc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// decodeGBK 解码中文目录名
func decodeGBK(input string) (string, error) {
	if notContainsChinese(input) {
		return input, nil
	}

	decoder := simplifiedchinese.GBK.NewDecoder()
	reader := transform.NewReader(strings.NewReader(input), decoder)
	decoded, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

func notContainsChinese(input string) bool {
	regex := regexp.MustCompile("[\u4e00-\u9fa5]") // 匹配中文字符的正则表达式
	return regex.MatchString(input)
}
