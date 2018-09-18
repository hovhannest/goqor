package config

import (
	"goqor1.0/app/Interface"
	logger "goqor1.0/logger/Interface"
)

type AppConfigurations struct {
}

func (appConfig *AppConfigurations) ConfigureApplication(app *Interface.AppConfig) {
	app.Config.SetDefault("log.level", logger.All)

	app.Config.SetDefault("server_port", 8080)
	app.Config.SetDefault("jwt_signing_method", "RS256")
	app.Config.SetDefault("use_https", false)
}

func New() Interface.MicroAppInterface {
	return &AppConfigurations{}
}