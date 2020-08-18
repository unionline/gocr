/*
@Time : 2020/4/14 23:36
@Author : Justin
@Description :
@File : basic_setting
@Software: GoLand
*/
package basic_setting_service

import (
	"gocr/models"
	"gocr/resposity"
)

var repo *resposity.BasicSettingRepo

func init() {
	repo = resposity.NewBasicSettingRepo()
}

func GetBasicSettingInfo(app_id, version string) (resp models.BasicSetting, err error) {

	item := models.NewBasicSetting()

	item.APPID = app_id
	item.Version = version

	resp, err = repo.GetBasicSettingInfo(item)
	return
}

func CreateBasicSetting(app_id, version, app_name, dev_log, about_me string) (err error) {

	var update_flag uint

	item := models.NewBasicSetting()

	item.APPID = app_id
	item.Version = version
	item.AppName = app_name
	item.DevLog = dev_log
	item.AboutMe = about_me
	item.UpdateFlag = update_flag

	err = repo.CreateBasicSetting(item)
	return
}
