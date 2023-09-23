package cmn

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

// 使用标准包进行Post请求，固定Content-Type:application/x-www-form-urlencoded，其他自定义headers格式为 K:V
func HttpPostForm(url string, formMap map[string]string, headers ...string) ([]byte, error) {

	sendBody := http.Request{}
	sendBody.ParseForm()

	for k, v := range formMap {
		sendBody.Form.Add(k, v)
	}
	sendData := sendBody.Form.Encode()

	client := &http.Client{}
	request, err := http.NewRequest("POST", url, strings.NewReader(sendData))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for i, max := 0, len(headers); i < max; i++ {
		strs := Split(headers[i], ":")
		if len(strs) > 1 {
			request.Header.Set(Trim(strs[0]), Trim(strs[1]))
		}
	}

	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(IntToString(res.StatusCode))
	}
	return io.ReadAll(res.Body)
}
