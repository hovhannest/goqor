package api

import (
	"github.com/go-chi/chi"
	"github.com/qor/render"
	"goqor1.0/app/Interface"
	"goqor1.0/app/auth"
	"goqor1.0/utils"
	"net/http"
)

type ApiConfigurations struct {
	Render *render.Render
}

func (appConfig *ApiConfigurations) ConfigureApplication(app *Interface.AppConfig) {
	appConfig.Render = utils.CopyRender(app.Render)
	app.Router.Route("/api/v1",  func(r chi.Router) {
		r.Route("/hello", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				if auth.App.Authority.Allow("admin", r) {
					w.Write([]byte("Authorised As Admin!! "))
				} else if auth.App.Authority.Allow("", r) {
					w.Write([]byte("Authorised "))
				} else {
					w.Write([]byte("Unauthorised"))
				}
			})
		})
	})
}

func New() Interface.MicroAppInterface {
	return &ApiConfigurations{}
}