package model

import "gorm.io/gorm"

type UserDetail struct {
	gorm.Model
	Pid    int    `gorm:"column:pid;comment:'pid';unique" json:"pid"`
	Tall   int32  `gorm:"column:tall;comment:'身高'" json:"tall"`
	Email  string `gorm:"column:email;comment:'email'" json:"email"`
	Sex    string `gorm:"column:sex;comment:'性别'" json:"sex"`
	Hobbit string `gorm:"column:hobbit;comment:'爱好'" json:"hobbit"`
}

func (u UserDetail) TableName() string {
	return "user_detail"
}
