package cmn

import (
	"io"
	"net/http"
	"time"
)

// 使用标准包进行Get请求，固定Content-Type:application/json;charset=UTF-8，其他自定义headers格式为 K:V
func HttpGetJson(url string, headers ...string) ([]byte, error) {

	// 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 请求头
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	for i, max := 0, len(headers); i < max; i++ {
		strs := Split(headers[i], ":")
		if len(strs) > 1 {
			req.Header.Set(Trim(strs[0]), Trim(Join(strs[1:], ":")))
		}
	}

	// 读取响应内容
	client := http.Client{Timeout: 1 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}

// 使用标准包进行Get请求，固定Content-Type:application/json;charset=UTF-8，其他自定义headers格式为 K:V
func HttpGetJsonTimeout(url string, timeout time.Duration, headers ...string) ([]byte, error) {

	// 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 请求头
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	for i, max := 0, len(headers); i < max; i++ {
		strs := Split(headers[i], ":")
		if len(strs) > 1 {
			req.Header.Set(Trim(strs[0]), Trim(Join(strs[1:], ":")))
		}
	}

	// 读取响应内容
	client := http.Client{Timeout: max(time.Second, min(timeout, 30*time.Second))}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}
