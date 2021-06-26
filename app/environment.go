package app

import (
	"os"
)

type Environment struct {
	App
	Cache Caching
	DB Database
}

func (env *Environment) Get() {
	env.EnvMode = os.Getenv("ENV_MODE")
	env.UseDB = os.Getenv("USE_DB")
	env.UseCacheDB = os.Getenv("USE_CACHE")
}

func (env *Environment) GetAPIPort() string {
	return os.Getenv(env.EnvMode + "_"+ "PORT")
}