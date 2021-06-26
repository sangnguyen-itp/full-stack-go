package database

import (
	"fmt"
	"full-stack-go/app"
	"gorm.io/gorm"
	"os"
)

type DBEnvironment struct {}

func (de *DBEnvironment) Connect(env *app.Environment, cfg *gorm.Config) {
	connectKey := fmt.Sprintf("%s_%s_", env.UseDB, env.EnvMode)
	if env.UseDB == "postgres" {
		env.DB = app.Database{
			Host : os.Getenv(connectKey + "HOST"),
			User:  os.Getenv(connectKey + "USER"),
			Password : os.Getenv(connectKey + "PASSWORD"),
			DbName : os.Getenv(connectKey + "DBNAME"),
			Port : os.Getenv(connectKey + "PORT"),
			SSL : os.Getenv(connectKey + "SSL"),
		}

		dbProvider := &PostgresDB{env:env}
		dbProvider.SetConnectionString()
		dbProvider.ConnectDB(NewBaseConfig(env), cfg)
	}
}

