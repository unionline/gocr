/*
@Time : 2020/4/14 14:39
@Author : FB
@File : ocr.go
@Software: GoLand
*/
package ocr_service

import (
	"gocr/env"
	"gocr/services/baidu_aip_service"
	"gocr/utils"
	"io/ioutil"
)

//https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic?
//url=https://c-ssl.duitang.com/uploads/item/201409/16/20140916141412_HAzPH.thumb.1900_0.jpeg
//&access_token=24.85158a93bc2e52ea1595d0ebdfe931ba.2592000.1574323037.282335-17593004

func getBasicReqUrl(image_type int) string {
	host := getReqUrlHost(image_type)
	access_token := getAccessToken()

	reqUrl := host + "?access_token=" + access_token + "&"

	return reqUrl
}

func getAccessToken() string {
	token := baidu_aip_service.GetAccessToken()
	return token
}

func getBaiduAIPHost() string {
	return "https://aip.baidubce.com/rest/2.0/ocr/v1/"

}

func getReqUrlHost(ocr_type int) string {

	baseUrl := getBaiduAIPHost()
	api := env.TypeMap[ocr_type]
	reqUrl := baseUrl + api

	return reqUrl
}

func getImageBase64Data(img_path string) (ret string, err error) {
	data_, err := ioutil.ReadFile(img_path)
	if err != nil {
		return
	}

	ret = utils.Base64EncodeByte(data_)
	return

}
