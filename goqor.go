package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/qor/render"
	"goqor1.0/app"
	"goqor1.0/app/Interface"
	"goqor1.0/app/auth"
	appConfig "goqor1.0/app/config"
	"goqor1.0/app/db"
	"goqor1.0/app/home"
	"goqor1.0/app/i18n"
	"goqor1.0/app/index"
	"goqor1.0/app/static"
	"goqor1.0/bindatafs"
	"goqor1.0/config"
	"goqor1.0/logger"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"
)

var App Interface.Application

func main() {
	config := config.Instance()
	l := logger.NewLogger(logger.NewLogConfig(&config))
	router := chi.NewRouter()
	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(30 * time.Second))

	assetFS := bindatafs.AssetFS
	assetPath := path.Join("app", "views")
	assetFS.RegisterPath(assetPath)
	assetFS.PrependPath(assetPath)

	render := render.New(&render.Config{})
	render.SetAssetFS(assetFS)

	AppConfig := app.AppConfig{
		Config: config,
		Logger: l,
		Router: router,
		Render: render,
	}

	App = app.Configure(&AppConfig)
	defer cleanup()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	signal.Notify(c, os.Interrupt, syscall.SIGKILL)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	App.Use(appConfig.New())
	App.Use(db.New())
	App.Use(i18n.New())
	App.Use(auth.New()) // call App.Use(db.New()) before this
	App.Use(static.New())
	App.Use(index.New())
	App.Use(home.New())

	//render.RegisterFuncMap("Greet", func(name string) string { return "Hello " + name })
	//render.RegisterFuncMap("render_header", func() template.HTML { return "<h1>render_header</h1>" })

	//render.RegisterViewPath(path.Join("app", "views", "themes", "theme1"))

	//router.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//	render.Execute("index", r.Context(), r, w)
	//})

	//router.Get("/home", func(w http.ResponseWriter, r *http.Request) {
	//	render.Layout("application18").Execute(path.Join("themes", "theme1", "home"), r.Context(), r, w)
	//})

	//funcMap := template.FuncMap{}

	//router.Get("/home111", func(w http.ResponseWriter, r *http.Request) {
	//	render.Layout("application18").Funcs(funcMap).Execute("home111", r.Context(), r, w)
	//})

	App.Run()
}

func cleanup() {
	App.ShutDown()
}