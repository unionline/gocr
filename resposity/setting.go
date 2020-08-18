/*
@Time : 2020/4/14 21:53
@Author : Justin
@Description :
@File : setting.go
@Software: GoLand
*/
package resposity

import (
	"gocr/initializers/db"
	"gocr/models"
)

type SettingRepo struct {
}

func NewSettingRepo() *SettingRepo {
	return &SettingRepo{}
}

func (*SettingRepo) GetSettingInfo(id uint) (setting models.Setting, err error) {
	err = db.Db.First(&setting, id).Error
	return
}

func (*SettingRepo) CreateSetting(setting *models.Setting) (err error) {
	err = Insert.CreateSetting(db.Db, setting)
	return
}
