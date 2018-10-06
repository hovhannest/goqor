package admin

import (
	"github.com/qor/admin"
	"goqor1.0/app/Interface"
	"goqor1.0/app/auth"
	"goqor1.0/app/db"
	"goqor1.0/app/i18n"
)

type AdminConfigurations struct {
	Coonfig admin.AdminConfig
}

func (appConfig *AdminConfigurations) ConfigureApplication(app *Interface.AppConfig) {
	admin := admin.New(&appConfig.Coonfig)
	app.Router.Mount("/admin", admin.NewServeMux("/admin"))
}

func New() Interface.MicroAppInterface {
	AdminConfigu := AdminConfigurations{}
	AdminConfigu.Coonfig = admin.AdminConfig{
		SiteName: "GoQor",
		DB: db.DB,
		Auth: auth.AdminAuth{},
		I18n: i18n.I18n,
	}
	return &AdminConfigu
}
