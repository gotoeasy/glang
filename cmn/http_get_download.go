package cmn

import (
	"io"
	"net/http"
	"os"
)

// 下载文件
func HttpDownload(url, saveAsPathFile string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	MkdirAll(Dir(saveAsPathFile))
	file, err := os.Create(saveAsPathFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}
