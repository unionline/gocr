/*
@Time : 2020/4/14 15:33
@Author : FB
@File : getResult.go
@Software: GoLand
*/
package ocr_service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"gocr/env"
	"gocr/initializers/config"
	"gocr/models"
	"gocr/services/baidu_aip_service"
	"gocr/services/ocr_log_service"
	"gocr/utils"
	"net/url"
	"strings"
)

func GetResultWithUrl(img_url string, image_type int, ip string) (ret string, err error) {

	reqUrl := GetReqUrlWithUrl(img_url, image_type)
	data, err := baidu_aip_service.GetOcrReqContent(reqUrl)
	if err != nil {
		return
	}

	result, err := parseToResult(data)
	if err != nil {
		return
	}

	var log models.OcrLog

	log.Type = "baidu-aip-ocr.GENERAL_BASIC.url"

	if result.ErrorCode != 0 {
		err = errors.New(result.ErrorMsg)
		log.Status = false

		// 重新生成Token
		if result.ErrorCode == 110 || result.ErrorCode == 111 {
			baidu_aip_service.GenAccessToken()
			log.Type = "baidu-aip-ocr.GENERAL_BASIC.url.ReGenAccessToken"
		}

	} else {
		ret, err = parseData(result, env.TYPE_OCR_GENERAL_BASIC)
		log.Status = true
		log.Result = data
		log.Content = ret
	}

	log.ImagePath = img_url
	log.Ip = ip

	//保留查询记录
	SaveLogOfOcr(&log)

	return
}

func GetResultWithImageUseUrl(img_url string, image_type int, ip string) (ret string, err error) {

	//获取img_url的image名称
	img_id := utils.GetImageId(img_url)
	ocr_log, err := ocr_log_service.GetOcrLogInfoByImageId(img_id)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	if ocr_log.ID > 0 {
		ret = ocr_log.Content
		return
	}

	img_url = config.Setting.Path.ServerAddress + strings.Replace(img_url, env.SAVE_PAHT_RES, env.REQUEST_URL_RES, 1)
	reqUrl := GetReqUrlWithUrl(img_url, image_type)
	data, err := baidu_aip_service.GetOcrReqContent(reqUrl)
	if err != nil {
		return
	}

	result, err := parseToResult(data)
	if err != nil {
		return
	}

	var log models.OcrLog

	log.Type = "baidu-aip-ocr.GENERAL_BASIC.image2url"

	if result.ErrorCode != 0 {
		err = errors.New(result.ErrorMsg)
		log.Status = false

		// 重新生成Token
		if result.ErrorCode == 110 || result.ErrorCode == 111 {
			baidu_aip_service.GenAccessToken()
			log.Type = "baidu-aip-ocr.GENERAL_BASIC.url.ReGenAccessToken"
		}

	} else {
		ret, err = parseData(result, env.TYPE_OCR_GENERAL_BASIC)
		log.Status = true
		log.Result = data
		log.Content = ret
		log.ImageID = img_id
	}

	log.ImagePath = img_url
	log.Ip = ip

	//保留查询记录
	SaveLogOfOcr(&log)

	return
}

func GetResultWithImagesUseUrl(img_url_arr []string, ip string) (ret string, err error) {
	ret_arr := make([]string, len(img_url_arr))
	for i, v := range img_url_arr {
		ret_arr[i], err = GetResultWithImageUseUrl(v, env.TYPE_OCR_GENERAL_BASIC, ip)
		if err != nil {
			return
		}
	}

	ret = strings.Join(ret_arr, "")

	return
}

func GetResultWithImage(img_path string, ip string) (ret string, err error) {

	reqUrl, err := GetReqUrlWithImage(img_path, env.TYPE_OCR_GENERAL_BASIC)
	data, err := baidu_aip_service.GetOcrReqContent(reqUrl)
	if err != nil {
		return
	}

	result, err := parseToResult(data)
	if err != nil {
		return
	}

	var log models.OcrLog

	if result.ErrorCode != 0 {
		err = errors.New(result.ErrorMsg)
		log.Status = false
	} else {
		ret, err = parseData(result, env.TYPE_OCR_GENERAL_BASIC)
		log.Status = true
		log.Result = data
		log.Content = ret
	}

	log.ImagePath = img_path
	log.Type = "baidu-aip-ocr.GENERAL_BASIC.image"
	log.Ip = ip

	//保留查询记录
	SaveLogOfOcr(&log)

	return
}

func GetResultWithImageBase64(data_base64 string, image_type int, ip string) (ret string, err error) {

	//reqUrl, err := GetReqUrlWithImage(img_path)
	img_url_escape := url.QueryEscape(data_base64)
	reqUrl := getBasicReqUrl(image_type) + "image=" + img_url_escape

	data, err := baidu_aip_service.GetOcrReqContent(reqUrl)
	if err != nil {
		return
	}

	result, err := parseToResult(data)
	if err != nil {
		return
	}

	var log models.OcrLog

	if result.ErrorCode != 0 {
		err = errors.New(result.ErrorMsg)
		log.Status = false
	} else {
		ret, err = parseData(result, env.TYPE_OCR_GENERAL_BASIC)
		log.Status = true
		log.Result = data
		log.Content = ret
	}

	log.Type = "test.baidu-aip-ocr.GENERAL_BASIC.image"
	log.Ip = ip

	//保留查询记录
	SaveLogOfOcr(&log)

	return
}
