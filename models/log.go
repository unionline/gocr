/*
@Time : 2020/4/21 15:30
@Author : FB
@File : log.go
@Software: GoLand
*/
package models

import (
	"gocr/initializers/config"
	"time"
)

// 操作的log
type Log struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	User      string    `json:"user" `   // 用户
	Action    string    `json:"action"`  // 行为
	Content   string    `json:"content"` // 内容
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (Log) TableName() string {
	return config.Setting.MySQL.DbPrefix + "log"
}

func NewLog() *Log {
	return &Log{}
}
