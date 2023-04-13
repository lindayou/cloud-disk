package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestNewTable(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:qwe123-=@tcp(172.16.15.166:3306)/cloud-disk?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		t.Fatal("Gorm New Engine Failed,err :", err)

	}
	err = db.AutoMigrate(UserBasic{}, UserDetail{})
	if err != nil {
		t.Fatal("Gorm New AutoMigrate Table Failed,err :", err)
	}

}
