package i18n

import (
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/database"
	"goqor1.0/app/Interface"
	"goqor1.0/app/db"
)

var I18n *i18n.I18n

type I18nConfigurations struct {
}

func (appConfig *I18nConfigurations) ConfigureApplication(app *Interface.AppConfig) {
	if(db.DB == nil) {
		panic("call App.Use(db.New()) before App.Use(i18n.New()). I don't have DB other ways.")
	}
	I18n = i18n.New(database.New(db.DB))
}

func New() Interface.MicroAppInterface {
	return &I18nConfigurations{
	}
}