/*
@Time : 2020/4/14 15:30
@Author : FB
@File : getReqUrl.go
@Software: GoLand
*/
package ocr_service

import (
	"net/url"
)

func GetReqUrlWithUrl(img_url string, image_type int) string {

	img_url_escape := url.QueryEscape(img_url)

	reqUrl := getBasicReqUrl(image_type) + "url=" + img_url_escape
	return reqUrl
}

func GetReqUrlWithImage(img_path string, image_type int) (reqUrl string, err error) {

	img_url, err := getImageBase64Data(img_path)
	if err != nil {
		return
	}

	img_url_escape := url.QueryEscape(img_url)
	reqUrl = getBasicReqUrl(image_type) + "image=" + img_url_escape

	return
}
