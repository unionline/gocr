/*
@Time : 2020/4/14 22:53
@Author : Justin
@Description :
@File : read.go
@Software: GoLand
*/
package crud

import (
	"github.com/jinzhu/gorm"
	"gocr/models"
)

type Read struct {
}

func (*Read) GetBasicSettingInfo(db *gorm.DB, input *models.BasicSetting) (output models.BasicSetting, err error) {
	err = db.First(&output, input).Error
	return
}

func (*Read) GetUserInfo(db *gorm.DB, id uint) (output models.User, err error) {
	err = db.First(&output, id).Error
	return
}

func (*Read) GetSettingInfo(db *gorm.DB, id uint) (output models.Setting, err error) {
	err = db.First(&output, id).Error
	return
}

func (*Read) GetOcrLogInfo(db *gorm.DB, id uint) (output models.OcrLog, err error) {
	err = db.First(&output, id).Error
	return
}

func (*Read) GetOcrLogInfoByImageId(db *gorm.DB, id string) (output models.OcrLog, err error) {
	err = db.First(&output, models.OcrLog{ImageID: id}).Error
	return
}

// Log
func (*Read) GetLogInfo(db *gorm.DB, id uint) (output models.Log, err error) {
	err = db.First(&output, id).Error
	return
}

// Image
func (*Read) GetImageInfo(db *gorm.DB, id uint) (output models.Image, err error) {
	err = db.First(&output, id).Error
	return
}

func (*Read) GetImageInfoByMd5(db *gorm.DB, md5 string) (output models.Image, err error) {
	err = db.Where(&models.Image{Md5: md5}).First(&output).Error
	return
}

func (*Read) GetImageList(db *gorm.DB) (output []models.Image, err error) {
	err = db.Find(&output).Error
	return
}

func (*Read) GetImageRecord(db *gorm.DB, xid string) (output models.Image, err error) {
	err = db.First(&output, models.Image{XID: xid}).Error
	return
}
