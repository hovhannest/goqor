package Interface

import (
	"github.com/go-chi/chi"
	"github.com/qor/render"
	config "goqor1.0/config/Interface"
	logger "goqor1.0/logger/Interface"
)

type AppConfig struct {
	Config config.Config
	Logger logger.Logger
	Router chi.Router
	Render *render.Render
}

// MicroAppInterface micro app interface
type MicroAppInterface interface {
	ConfigureApplication(*AppConfig)
}

type Application interface {
	GetConfig() *AppConfig
	SetConfig(*AppConfig)

	Run()
	ShutDown()

	Use(MicroAppInterface)
}

var App Application
