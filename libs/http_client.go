package libs

import (
	"bytes"
	"fmt"
	"github.com/tal-tech/go-zero/core/jsonx"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	smsConfig Sms
}

type Resp struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{}
}


func (hc *HttpClient) SetConfig(config Sms){
	hc.smsConfig = config

	fmt.Println(config)
}

func (hc *HttpClient) HttpPost(uri string, request interface{}) (interface{},error) {
	 url := hc.smsConfig.Hostname + uri
	 params, err := jsonx.Marshal(request)
	 if err != nil {
	 	fmt.Println("err fail")
		 return "", nil
	 }

	reader := bytes.NewReader(params)
	client := &http.Client{}
	requests, err := http.NewRequest("POST", url, reader)
	if err != nil {
		// handle error
	}

	requests.Header.Set("Content-Type", "application/json")

	response,err := client.Do(requests)
	defer response.Body.Close()
	body,err := ioutil.ReadAll(response.Body)
	rep :=Resp{}
	jsonx.Unmarshal(body,&rep)
	var result interface{}
	if rep.Code == 200 {
		result = rep.Data
	}

	return result,nil
}