package app

import (
	"github.com/qor/render"
	"github.com/qor/session/manager"
	"goqor1.0/app/Interface"
	"html/template"
	"net/http"
	"fmt"
)

type AppConfig = Interface.AppConfig


type Application struct {
	Config AppConfig
	AppName string
	configFile string
	configFormat string
}

type AppGoqor struct {

}

func init() {
	Interface.App = &Application{
		AppName: "GoQor",
		configFile: "config",
		configFormat: "json",
	}
}

func Configure(appConfig *AppConfig) Interface.Application {
	Interface.App.SetConfig(appConfig)
	app := Interface.App.(*Application)
	app.Config.Config.Read(app.AppName, app.configFile, app.configFormat)
	app.configureRender()
	return Interface.App
}

func (app *Application) GetConfig() *Interface.AppConfig {
	return &app.Config
}

func (app *Application) SetConfig(config *Interface.AppConfig) {
	app.Config = *config
}

func (app *Application) Run() {
	port := app.Config.Config.GetString("server_port")
	addr := ":" + port
	fmt.Printf("Running server in localhost%s\n", addr)
	http.ListenAndServe(addr, manager.SessionManager.Middleware(Interface.App.GetConfig().Router))
}

func (app *Application) ShutDown() {
	app.Config.Config.Save(app.configFile + "." + app.configFormat)
}

func (app *Application) Use(microApp Interface.MicroAppInterface) {
	microApp.ConfigureApplication(&app.Config)
}


func (app *Application) configureRender() {
	r := app.Config.Render
	oldFuncMapMaker := r.FuncMapMaker
	app.Config.Render.FuncMapMaker = func(render *render.Render, req *http.Request, w http.ResponseWriter) template.FuncMap {
		funcMap := template.FuncMap{}
		if oldFuncMapMaker != nil {
			funcMap = oldFuncMapMaker(render, req, w)
		}

		funcMap["Greet"] = func(name string) string { return "Hello " + name }
		funcMap["render_header"] = func() template.HTML { return "<h1>render_header</h1>" }

		return funcMap
	}
}