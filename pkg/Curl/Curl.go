package Curl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Get(getUrl string, getData map[string]string) (resp string, err error)  {

	params := url.Values{}
	Url, _:= url.Parse(getUrl)
	if len(getData) > 0 {
		for k,v := range getData {
			params.Set(k,v)
		}
	}
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
	response,_ := http.Get(urlPath)

	body, _ := ioutil.ReadAll(response.Body)

	return string(body),nil
}

func Post(postUrl string, postData map[string]string, headers map[string]string) (resp string, err error) {

	// 取地址方式实例化client
	client := &http.Client{}

	// 组装请求体
	var data = url.Values{}
	if len(postData) > 0 {
		for k,v := range postData {
			data.Add(k, v)
		}
	}
	request, err := http.NewRequest("POST", postUrl, strings.NewReader(data.Encode()))

	// 组装请求头
	if len(headers) > 0 {
		for k,v := range headers {
			request.Header.Set(k,v)
		}
	}

	response, _ := client.Do(request)
	body, _ := ioutil.ReadAll(response.Body)

	return string(body),nil
}