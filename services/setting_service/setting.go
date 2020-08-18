/*
@Time : 2020/4/14 23:36
@Author : Justin
@Description :
@File : setting
@Software: GoLand
*/
package setting_service

import (
	"gocr/models"
	"gocr/resposity"
)

// SettingService ...
type SettingService struct {
	Setting *models.Setting
}

// NewSettingService ...
func NewSettingService() *SettingService {
	return &SettingService{
		Setting: models.NewSetting(),
	}
}

var userRepo *resposity.SettingRepo

func init() {
	userRepo = resposity.NewSettingRepo()
}

// GetSettingInfo ...
func (service *SettingService) GetSettingInfo(id uint) (resp models.Setting, err error) {
	resp, err = userRepo.GetSettingInfo(id)
	return
}

// GetSettingInfo ...
func (service *SettingService) CreateSetting(app_id, app_key, app_secret, type_setting string) (err error) {

	item := models.NewSetting()
	item.AppId = app_id
	item.AppKey = app_key
	item.AppSecret = app_secret
	item.Type = type_setting

	err = userRepo.CreateSetting(item)
	return
}
