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
		host := os.Getenv(connectKey + "HOST")
		user := os.Getenv(connectKey + "USER")
		password := os.Getenv(connectKey + "PASSWORD")
		dbName := os.Getenv(connectKey + "DBNAME")
		port := os.Getenv(connectKey + "PORT")
		sslMode := os.Getenv(connectKey + "SSL")

		dbProvider := new(PostgresDB)
		dbProvider.SetConnectionString(host, user, password, dbName, port, sslMode)
		dbProvider.ConnectDB(NewBaseConfig(env), cfg)
	}
}

