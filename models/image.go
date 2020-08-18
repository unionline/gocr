/*
@Time : 2020/4/16 10:15
@Author : FB
@File : image.go
@Software: GoLand
*/
package models

import (
	"gocr/initializers/config"
	"time"
)

type Image struct {
	ID            uint      `gorm:"primary_key" json:"id"`
	XID           string    `json:"xid"`                        // XID
	Name          string    `json:"name"`                       // 原始的图片名称
	Size          string    `json:"size"`                       // 原始的图片大小保留两位小数
	FileName      string    `json:"file_name"`                  // 数据库图片名称
	ThumbnailName string    `json:"thumbnail_name"`             // 数据库图片缩略图名称
	From          string    `json:"from" gorm:"default:'gocr'"` // 图片来源的平台app_id,有gocr，gocr2，pyocr
	Md5           string    `json:"md5"`                        // MD5
	Etc           string    `json:"etc"`                        // 备注
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (Image) TableName() string {
	return config.Setting.MySQL.DbPrefix + "image"
}

func NewImage() *Image {
	return &Image{}
}
