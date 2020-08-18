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

type BasicSettingRepo struct {
}

func NewBasicSettingRepo() *BasicSettingRepo {
	return &BasicSettingRepo{}
}

func (*BasicSettingRepo) GetBasicSettingInfo(basicSetting *models.BasicSetting) (item models.BasicSetting, err error) {
	item, err = Query.GetBasicSettingInfo(db.Db, basicSetting)
	return
}

func (*BasicSettingRepo) CreateBasicSetting(basicSetting *models.BasicSetting) (err error) {
	err = Insert.CreateBasicSetting(db.Db, basicSetting)
	return
}
