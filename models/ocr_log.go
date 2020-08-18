/*
@Time : 2020/4/14 21:48
@Author : Justin
@Description :
@File : log.go
@Software: GoLand
*/
package models

import (
	"gocr/initializers/config"
	"time"
)

type OcrLog struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	ImageID   string    `json:"image_id"`                            // 图片名称
	Result    string    `json:"result" type:varchar(2048)`           // 请求识别原始结果
	Content   string    `json:"content" type:varchar(2048)`          // 请求识别结果
	ImagePath string    `json:"image_path"`                          // 网络图片地址或者本地图片访问地址
	Status    bool      `json:"status"`                              // 请求识别结果状态
	Type      string    `json:"type" gorm:"default:'baidu-aip-ocr'"` // 默认是baidu-aip-ocr
	Ip        string    `json:"ip"`                                  // 访问请求的IP地址
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (OcrLog) TableName() string {
	return config.Setting.MySQL.DbPrefix + "ocr_log"
}

func NewOcrLog() *OcrLog {
	return &OcrLog{}
}
