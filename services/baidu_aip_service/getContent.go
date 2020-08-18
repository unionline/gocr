/*
@Time : 2020/4/14 16:16
@Author : FB
@File : req.go
@Software: GoLand
*/
package baidu_aip_service

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func GetContent(req_url string) (ret []byte, err error) {

	// 增加超时
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		MaxIdleConns: 2,
	}
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}

	fmt.Println("req_url=", req_url)
	resp, err := client.Get(req_url)

	defer resp.Body.Close()

	ret, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	fmt.Println("ret=", string(ret))

	return
}

func GetOcrReqContent(req_url string) (ret string, err error) {

	// 增加超时
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		MaxIdleConns: 2,
	}
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}

	request, err := http.NewRequest("POST", req_url, nil)
	//增加header选项
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	if err != nil {
		panic(err)
	}

	fmt.Println("req_url=", req_url)
	resp, err := client.Do(request)

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	ret = string(data)
	fmt.Println("ret=", ret)

	return
}
