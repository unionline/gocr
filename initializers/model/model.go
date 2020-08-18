/*
@Time : 2020/3/12 16:50
@Author : FB
@File : model.go
@Software: GoLand
*/
package model

import (
	"gocr/initializers/db"
	"gocr/models"
)

func Init() {

	if !db.Db.HasTable(models.User.TableName(models.User{})) {
		db.Db.CreateTable(models.User{})
	}
	db.Db.AutoMigrate(models.User{})

	if !db.Db.HasTable(models.Setting.TableName(models.Setting{})) {
		db.Db.CreateTable(models.Setting{})
	}
	db.Db.AutoMigrate(models.Setting{})

	if !db.Db.HasTable(models.OcrLog.TableName(models.OcrLog{})) {
		db.Db.CreateTable(models.OcrLog{})
	}
	db.Db.AutoMigrate(models.OcrLog{})

	if !db.Db.HasTable(models.Image.TableName(models.Image{})) {
		db.Db.CreateTable(models.Image{})
	}
	db.Db.AutoMigrate(models.Image{})

	if !db.Db.HasTable(models.BasicSetting.TableName(models.BasicSetting{})) {
		db.Db.CreateTable(models.BasicSetting{})
	}
	db.Db.AutoMigrate(models.BasicSetting{})

	if !db.Db.HasTable(models.Log.TableName(models.Log{})) {
		db.Db.CreateTable(models.Log{})
	}
	db.Db.AutoMigrate(models.Log{})

}
