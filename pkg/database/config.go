package database

import (
	"full-stack-go/app"
	"full-stack-go/constants"
)

type BaseConfig struct {
	IsDebug bool
	IsMigration bool
}

func NewBaseConfig(env *app.Environment) BaseConfig {
	baseCfg := &BaseConfig{}
	baseCfg.getBaseConfig(env)
	return *baseCfg
}

func (b *BaseConfig) getBaseConfig(env *app.Environment) {
	switch env.EnvMode {
	case constants.DEV:
		b.IsDebug = true
		b.IsMigration = true
	case constants.TEST:
		b.IsDebug = false
		b.IsMigration = true
	case constants.PRODUCTION:
		b.IsDebug = false
		b.IsMigration = true
	default:
		b.IsDebug = true
		b.IsMigration = true
	}
}

