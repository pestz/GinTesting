package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username  string
	Password  string
	Nickname  string
	Status    string
	Telephone string
	Avatar    string `gorm:"size:1000"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)
