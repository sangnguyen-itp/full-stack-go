package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	connectionString string
}

func (p *PostgresDB) ConnectDB(baseCfg BaseConfig, config *gorm.Config) {
	if p.connectionString != "" {
		db, err := gorm.Open(postgres.Open(p.connectionString), config)
		if err != nil {
			panic(err)
		}
		if baseCfg.IsDebug {
			DBInstance = db.Debug()
		} else {
			DBInstance = db
		}
		if baseCfg.IsMigration {
			err := MigrationTables()
			if err != nil {
				fmt.Println("failed")
				panic(err)
			}
		}
		return
	}
	panic("error: your connection string is invalid")
}

func (p *PostgresDB) SetConnectionString(
	host, user, password, dbname, port string,
	sslMode string,
	){
	p.connectionString = fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		user,
		password,
		host,
		port,
		dbname,
		sslMode)
}

var _ Connection = &PostgresDB{}

