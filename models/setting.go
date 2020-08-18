/*
@Time : 2020/4/14 21:53
@Author : Justin
@Description :
@File : setting.go
@Software: GoLand
*/
package models

import (
	"gocr/initializers/config"
	"time"
)

type Setting struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	AppId     string    `json:"app_id"`
	AppKey    string    `json:"app_key"`
	AppSecret string    `json:"app_secret"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Setting) TableName() string {

	return config.Setting.MySQL.DbPrefix + "setting"
}

func NewSetting() *Setting {
	return &Setting{}
}
