/*
@Time : 2020/4/14 16:35
@Author : FB
@File : getAppInfo.go
@Software: GoLand
*/
package baidu_aip_service

import (
	"gocr/models"
	"gocr/resposity"
)

func getAppInfo() (ret models.Setting, err error) {

	ret, err = resposity.NewSettingRepo().GetSettingInfo(1)
	if err != nil {
		return
	}

	return
}
