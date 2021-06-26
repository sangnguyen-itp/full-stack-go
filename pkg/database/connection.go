package database

import (
	"gorm.io/gorm"
)

type Connection interface {
	CreateDatabase(dbName string) error
	GetDbName() string
	ConnectDB(basicConfig BaseConfig, config *gorm.Config)
	SetConnectionString()
}

var DBInstance *gorm.DB

func GetDBInstance() *gorm.DB {
	return DBInstance
}


