/*
@Time : 2020/4/17 10:46
@Author : FB
@File : basic_setting.go
@Software: GoLand
*/
package models

import (
	"gocr/initializers/config"
	"time"
)

type BasicSetting struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	APPID      string    `json:"app_id"`      // 软件ID
	Version    string    `json:"version"`     // 软件版本
	AppName    string    `json:"app_name"`    // 软件名称
	DevLog     string    `json:"dev_log"`     // 开发日志，高版本包括低版本开发日志
	AboutMe    string    `json:"about_me"`    // 关于我
	UpdateFlag uint      `json:"update_flag"` // 0,1,2 =>不提示更新，提示更新，强制更新
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (BasicSetting) TableName() string {

	return config.Setting.MySQL.DbPrefix + "basic_setting"
}

func NewBasicSetting() *BasicSetting {
	return &BasicSetting{}
}
