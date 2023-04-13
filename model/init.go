package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Init(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Gorm New Engine Failed,err :", err)
		return nil
	}
	return db

}
