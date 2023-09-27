package cmn

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// 使用标准包进行Post请求，固定Content-Type:multipart/form-data，其他自定义headers格式为 K:V
func HttpUploadFile(url string, filePath string, headers ...string) ([]byte, error) {
	// 打开要上传的文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 创建一个缓冲区，用于构建multipart请求体
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加文件到请求体
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	// 关闭multipart写入器
	writer.Close()

	// 发起POST请求
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	// req.Header.Set("Content-Type", "multipart/form-data")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	for i, max := 0, len(headers); i < max; i++ {
		strs := Split(headers[i], ":")
		if len(strs) > 1 {
			req.Header.Set(Trim(strs[0]), Trim(strs[1]))
		}
	}

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// 处理响应
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(IntToString(res.StatusCode))
	}
	return io.ReadAll(res.Body)
}
