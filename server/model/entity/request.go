package entity

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Request struct {
	Method  string
	URL     string
	Params  map[string]string
	Headers map[string]string
	Body    interface{}
}

func NewRequest(method, url string) *Request {
	return &Request{
		Method:  method,
		URL:     url,
		Params:  make(map[string]string),
		Headers: make(map[string]string),
	}
}
func (req *Request) SetParam(key, value string) {
	req.Params[key] = value
}
func (req *Request) SetHeader(key, value string) {
	req.Headers[key] = value
}
func (req *Request) SetJsonBody(data interface{}) {
	req.Body = data
}
func (req *Request) SetFormBody(data map[string]string) {
	req.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	req.Body = data
}
func (req *Request) Send() (*http.Response, error) {
	// 构建 URL
	reqURL := req.URL
	if len(req.Params) > 0 {
		params := url.Values{}
		for key, value := range req.Params {
			params.Add(key, value)
		}
		reqURL += "?" + params.Encode()
	}

	// 构建请求
	var bodyReader io.Reader
	if req.Body != nil {
		switch v := req.Body.(type) {
		case string:
			bodyReader = strings.NewReader(v)
		case url.Values:
			bodyReader = strings.NewReader(v.Encode())
		case []byte:
			bodyReader = bytes.NewReader(v)
		default:
			return nil, fmt.Errorf("Unsupported request body type")
		}
	}

	request, err := http.NewRequest(req.Method, reqURL, bodyReader)
	if err != nil {
		return nil, err
	}

	for key, value := range req.Headers {
		request.Header.Set(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
