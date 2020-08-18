package models

import (
	"gocr/initializers/config"
	"time"
)

// User ...
type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `json:"username"`
	Nickname  string    `json:"nickname"`
	Password  string    `json:"password"`
	Salt      string    `json:"salt"`
	Phone     string    `json:"phone"`
	Mail      string    `json:"mail"`
	LastLogin time.Time `json:"last_login"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {

	return config.Setting.MySQL.DbPrefix + "user"
}

// NewUser ...
func NewUser() *User {
	return &User{}
}
