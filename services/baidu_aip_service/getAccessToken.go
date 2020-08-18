/*
@Time : 2020/4/14 16:31
@Author : FB
@File : getAccessToken.go
@Software: GoLand
*/
package baidu_aip_service

import (
	"encoding/json"
	"gocr/entities"
	"gocr/initializers/config"
	"gocr/models"
	"gocr/services/log_service"
	"gocr/utils"
	"io/ioutil"
	"log"
)

var EXPIRES_IN float64 = 2592000 //一个月
var ACCESS_TOKEN = ""

func GetAccessToken() string {

	if ACCESS_TOKEN == "" {
		data, err := readTokenInfo()
		if err != nil {
			err = GenAccessToken()
			if err != nil {
				log.Println("err=", err.Error())
				return ""
			}
			return ACCESS_TOKEN

		}

		// TODO 临时代码：尝请求一次来验证Token是否过期

		return data.AccessToken
	}

	return ACCESS_TOKEN

}

func testTokenValidByPing() {

}

func readTokenInfo() (ret entities.TokenResp, err error) {

	data, err := ioutil.ReadFile(config.Setting.Path.PathTokenFile)
	if err != nil {
		return
	}

	json.Unmarshal(data, &ret)

	return

}

func saveTokenInfo(input []byte) (err error) {

	err = utils.WriterFileConent(config.Setting.Path.PathTokenFile, input)
	if err != nil {
		return
	}

	return
}

func GenAccessToken() (err error) {

	tokenInfo, err := getOauthInfoByte()
	if err != nil {
		return
	}

	var data entities.TokenResp
	err = json.Unmarshal(tokenInfo, &data)
	if err != nil {
		return
	}

	EXPIRES_IN = data.ExpiresIn
	ACCESS_TOKEN = data.AccessToken

	// 保存TokenInfo记录到本地
	err = saveTokenInfo(tokenInfo)
	if err != nil {
		return
	}

	var item_log models.Log
	item_log.User = "admin"
	item_log.Action = "GenAccessToken"
	item_log.Content = ACCESS_TOKEN

	err = log_service.CreateLog(&item_log)
	if err != nil {
		return
	}

	return
}

func getOauthInfoByte() (ret []byte, err error) {

	info, err := getAppInfo()
	apiKey := info.AppKey
	secretKey := info.AppSecret

	reqUrl := getOauthReqUrl() + "client_id=" + apiKey + "&client_secret=" + secretKey

	ret, err = GetContent(reqUrl)
	if err != nil {
		return
	}

	return
}

func getOauthInfo() (ret entities.TokenResp, err error) {
	data, err := getOauthInfoByte()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &ret)
	if err != nil {
		return
	}

	return
}

func getOauthReqUrl() string {
	ret := "https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&"
	return ret
}
