package main

import (
	"full-stack-go/app"
	"full-stack-go/pkg/caching"
	"full-stack-go/pkg/database"
	"full-stack-go/routes"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"os"
	"strconv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	// initialize env
	env := new(app.Environment)
	env.Get()

	// Database
	dbEnv := new(database.DBEnvironment)
	isDefaultConstraint, err := strconv.ParseBool(os.Getenv("DISABLE_CONSTRAINTS"))
	if err != nil {
		panic(err)
	}
	dbEnv.Connect(env, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: isDefaultConstraint,
	})

	// routes
	r := new(routes.Router)
	r.SetupRouter(env)

	// Seed data

	// Caching
	caching.NewCacheService(env)

	// Third parties

	// Background Job

	r.Run()
}
