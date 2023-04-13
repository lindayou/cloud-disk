package model

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Id       int    `gorm:"column:id;primary_key" json:"id"`
	Name     string `gorm:"column:name;comment:'用户名'" json:"name"`
	Password string `gorm:"column:password;comment:'密码'" json:"password"`
}

func (UserBasic) TableName() string {
	return "user_basic"
}
