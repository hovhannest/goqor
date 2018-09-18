package static

import (
	"github.com/go-chi/chi"
	"github.com/qor/render"
	"goqor1.0/app/Interface"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type StaticConfigurations struct {
	Render *render.Render
}

func (appConfig *StaticConfigurations) ConfigureApplication(app *Interface.AppConfig) {
	app.Router.Route("/assets", func(r chi.Router) {
		workDir, _ := os.Getwd()
		filesDir := filepath.Join(workDir, "app", "static") // path is included
		FileServer(r, "/", http.Dir(filesDir))

		//r.Get("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//	w.Write([]byte("hi fs"))
		//}))
	})
}

func New() Interface.MicroAppInterface {
	return &StaticConfigurations{}
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
