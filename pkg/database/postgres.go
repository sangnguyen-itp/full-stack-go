package database

import (
	"bytes"
	"fmt"
	"full-stack-go/app"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"os/exec"
)

type PostgresDB struct {
	connectionString string
	env *app.Environment
}

func (p *PostgresDB) CreateDatabase(dbName string) error {
	cmd := exec.Command("createdb", "-p", p.env.DB.Port, "-h", p.env.DB.Host, "-U", p.env.DB.User, "-e", p.env.DB.DbName)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	log.Printf("Output: %q\n", out.String())
	return nil
}

func (p *PostgresDB) GetDbName() string {
	envKey := p.env.UseCacheDB + "_" + p.env.EnvMode + "_"
	return os.Getenv(envKey + "DBNAME")
}

func (p *PostgresDB) ConnectDB(baseCfg BaseConfig, config *gorm.Config) {
	if p.connectionString != "" {
		// create database if it's not exist
		p.CreateDatabase(p.GetDbName())
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

func (p *PostgresDB) SetConnectionString(){
	p.connectionString = fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		p.env.DB.User,
		p.env.DB.Password,
		p.env.DB.Host,
		p.env.DB.Port,
		p.env.DB.DbName,
		p.env.DB.SSL)
}

var _ Connection = &PostgresDB{}

