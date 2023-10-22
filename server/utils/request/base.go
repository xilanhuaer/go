package request

import (
	"fmt"
	"interface/global"
	"io"
	"net/http"
	"net/url"
)

func NewRequest(path string) {
	apiUrl := global.Host + path
	data := url.Values{}
	data.Set("page", "1")
	data.Set("page_size", "10")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}
	u.RawQuery = data.Encode()
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
