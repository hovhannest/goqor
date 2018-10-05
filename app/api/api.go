package api

import (
	"github.com/go-chi/chi"
	"github.com/qor/render"
	"goqor1.0/app/Interface"
	"goqor1.0/utils"
)

type ApiConfigurations struct {
	Render *render.Render
}

func (appConfig *ApiConfigurations) ConfigureApplication(app *Interface.AppConfig) {
	appConfig.Render = utils.CopyRender(app.Render)
	app.Router.Route("/api/v1",  func(r chi.Router) {
	})
}

func New() Interface.MicroAppInterface {
	return &ApiConfigurations{}
}