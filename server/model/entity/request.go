package entity

import (
	"io"
	"net/http"
)

type Request struct {
	Url      string         `json:"url"`
	Method   string         `json:"method"`
	Params   string         `json:"params"`
	Headers  string         `json:"headers"`
	Body     io.ReadCloser  `json:"body"`
	Response *http.Response `json:"response"`
}
