package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"go/types"
)

type shard map[string]int
type hit map[string]interface{}
type aaa []map[string]interface{}
type Code struct {
	Took int `json:took`
	TimeOut bool `json:time_out`
	_Shards shard `json:shards`
	Hits   hit 	`json:hits`
	}


func main(){
	url := "http://10.1.0.12:9200/test-2018-06/zabbix/_search"
	request, err := http.Get( url)
	if err != nil{
		fmt.Println(err)
	}
	respBytes, err := ioutil.ReadAll(request.Body)
	if err != nil{
		fmt.Println(err)
	}
	ParserIdJson(respBytes)
}


func ParserIdJson(body []byte)(id string,err error) {
	var code Code
	err = json.Unmarshal(body, &code)
	if err != nil {
		return "", nil
	}
	switch v := code.Hits["hits"].(type) {
	case types.Slice:
		fmt.Println(v.Elem().String())
	default:
		fmt.Println(v)
	}

	return "", nil
}