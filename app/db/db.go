package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goqor1.0/app/Interface"
	"goqor1.0/config"
	"goqor1.0/logger"
)

var (
	_config config.Config
	_packageName = "db"
	log logger.Logger
	DB *gorm.DB
)

func keyForPackage(key string) string {
	return  _packageName + "." + key
}

type Db struct {

}

func (appConfig *Db) ConfigureApplication(app *Interface.AppConfig) {
	log = app.Logger
	_config = app.Config

	_config.SetDefault(keyForPackage("adapter"),"mysql")
	_config.SetDefault(keyForPackage("User"),"coder")
	_config.SetDefault(keyForPackage("Password"),"abc123")
	_config.SetDefault(keyForPackage("Host"),"192.168.254.130")
	_config.SetDefault(keyForPackage("Port"),"3306")
	_config.SetDefault(keyForPackage("dbName"),"goqor")
	_config.SetDefault(keyForPackage("DEBUG"),true)

	var err error

	DB, err = gorm.Open(_config.GetString(keyForPackage("adapter")), fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local",
		_config.GetString(keyForPackage("User")),
		_config.GetString(keyForPackage("Password")),
		_config.GetString(keyForPackage("Host")),
		_config.GetString(keyForPackage("Port")),
		_config.GetString(keyForPackage("dbName"))))

	// DB = DB.Set("gorm:table_options", "CHARSET=utf8")
	if err == nil {
		if _config.GetBool(keyForPackage("DEBUG")) {
			DB.LogMode(true)
		}
		log.Debug("DB connected!")

		//l10n.RegisterCallbacks(DB)
		//sorting.RegisterCallbacks(DB)
		//validations.RegisterCallbacks(DB)
		//media.RegisterCallbacks(DB)
		//publish2.RegisterCallbacks(DB)

		MigrateAll()
	} else {
		panic(err)
	}
}

func New() Interface.MicroAppInterface {
	return &Db{}
}
