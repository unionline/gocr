/*
@Time : 2020/4/14 22:53
@Author : Justin
@Description :
@File : create.go
@Software: GoLand
*/
package crud

import (
	"github.com/jinzhu/gorm"
	"gocr/models"
)

type Insert struct {
}

func (*Insert) Create(db *gorm.DB, model interface{}) (err error) {
	err = db.Create(&model).Error
	return
}

func (*Insert) CreateBasicSetting(db *gorm.DB, model *models.BasicSetting) (err error) {
	err = db.Create(&model).Error
	return
}

func (*Insert) CreateSetting(db *gorm.DB, model *models.Setting) (err error) {
	err = db.Create(&model).Error
	return
}

func (*Insert) CreateLog(db *gorm.DB, model *models.Log) (err error) {
	err = db.Create(&model).Error
	return
}

func (*Insert) CreateOcrLog(db *gorm.DB, model *models.OcrLog) (err error) {
	err = db.Create(&model).Error
	return
}

func (*Insert) CreateUser(db *gorm.DB, model *models.User) (err error) {
	err = db.Create(&model).Error
	return
}

func (*Insert) CreateImage(db *gorm.DB, model *models.Image) (err error) {
	err = db.Create(&model).Error
	return
}
