package svc

import (
	"cloud-disk/core/internal/config"
	"cloud-disk/model"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     model.Init(c.Mysql.DataSource),
	}
}
