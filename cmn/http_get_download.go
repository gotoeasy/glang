package cmn

import (
	"io"
	"net/http"
	"os"
)

// 下载文件，自定义headers格式为 K:V
func HttpDownload(url, saveAsPathFile string, headers ...string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// 请求头
	for i, max := 0, len(headers); i < max; i++ {
		strs := Split(headers[i], ":")
		if len(strs) > 1 {
			req.Header.Set(Trim(strs[0]), Trim(Join(strs[1:], ":")))
		}
	}

	response, err := client.Do(req)
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
