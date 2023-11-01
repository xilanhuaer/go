package json

import (
	"encoding/json"

	"github.com/oliveagle/jsonpath"
)

type JsonPathService struct {
}

func (jsonPathService *JsonPathService) ReadJson(data, path string) (interface{}, error) {
	var jsonData interface{}
	err := json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		return nil, err
	}
	res, err := jsonpath.JsonPathLookup(jsonData, path)
	if err != nil {
		return nil, err
	}
	return res, nil

}
