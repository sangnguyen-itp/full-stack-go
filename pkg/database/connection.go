package database

import (
	"gorm.io/gorm"
)

type Connection interface {
	ConnectDB(basicConfig BaseConfig, config *gorm.Config)
	SetConnectionString(
		host, user, password, dbname, port, sslMode string,
	)
}

var DBInstance *gorm.DB

func GetDBInstance() *gorm.DB {
	return DBInstance
}


