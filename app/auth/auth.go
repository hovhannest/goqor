package auth

import (
	"github.com/qor/auth"
	"github.com/qor/auth/providers/password"
	"github.com/qor/auth_themes/clean"
	"github.com/qor/qor"
	"github.com/qor/qor/utils"
	"github.com/qor/redirect_back"
	"github.com/qor/render"
	"github.com/qor/session/manager"
	"goqor1.0/app/Interface"
	"goqor1.0/app/db"
	"goqor1.0/app/i18n"
	u "goqor1.0/utils"
	"html/template"
	"net/http"
)

type AuthConfigurations struct {
	Auth *auth.Auth
}

func (appConfig *AuthConfigurations) ConfigureApplication(app *Interface.AppConfig) {
	if(appConfig.Auth == nil) {
		if(db.DB == nil) {
			panic("call App.Use(db.New()) before App.Use(auth.New()). I don't have DB other ways.")
		}
		if(i18n.I18n == nil) {
			panic("call App.Use(i18n.New()) before App.Use(auth.New()). I don't have i18n other ways.")
		}
		appConfig.Auth =  clean.New(&auth.Config{
			DB: db.DB,
			Render: u.CopyRender(app.Render, func(render *render.Render, req *http.Request, w http.ResponseWriter) template.FuncMap {
				return template.FuncMap{
					"t": func(key string, args ...interface{}) template.HTML {
						return i18n.I18n.T(utils.GetLocale(&qor.Context{Request: req}), key, args...)
					},
				}}, "application18"),
			Redirector: auth.Redirector{
				RedirectBack: redirect_back.New(&redirect_back.Config{
					SessionManager:  manager.SessionManager,
					IgnoredPrefixes: []string{"/auth"},
				},
			)},
		})
	}

	appConfig.Auth.RegisterProvider(password.New(&password.Config{}))

	app.Router.Mount("/auth/", appConfig.Auth.NewServeMux())
}

func New() Interface.MicroAppInterface {
	return &AuthConfigurations{
		Auth: nil,
	}
}
