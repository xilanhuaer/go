package impl

import (
	"io"
	"net/http"
)

type RequestService struct {
}

// Request 发送请求
func (r *RequestService) Request(method, url string, params, headers map[string]interface{}, body io.Reader) (status int, resBody string, err error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return 0, "", err
	}
	if params != nil {
		query := request.URL.Query()
		for k, v := range params {
			query.Add(k, v.(string))
		}
		request.URL.RawQuery = query.Encode()
	}
	if headers != nil {
		for k, v := range headers {
			request.Header.Add(k, v.(string))
		}
	}
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return 0, "", err
	}
	defer res.Body.Close()
	status = res.StatusCode
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, "", err
	}
	resBody = string(b)
	return status, resBody, nil
}
