package index

import (
	"github.com/go-chi/chi"
	"github.com/qor/render"
	"goqor1.0/app/Interface"
	"goqor1.0/app/auth"
	"goqor1.0/utils"
	"net/http"
	"path"
)

type IngexConfigurations struct {
	Render *render.Render
}

func (appConfig *IngexConfigurations) ConfigureApplication(app *Interface.AppConfig) {
	appConfig.Render = utils.CopyRender(app.Render)
	appConfig.Render.RegisterViewPath(path.Join("app", "index", "views"))
	app.Router.Route("/", func(r chi.Router) {
		r.Get("/", appConfig.index)
		r.Route("/hello", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla pulvinar urna quis porttitor viverra. Fusce quis nisi orci. Vestibulum a justo commodo, vestibulum enim vitae, venenatis erat. Donec aliquet ligula sit amet velit posuere scelerisque. Suspendisse gravida felis a arcu pharetra dictum. Nunc egestas dictum nibh, sed euismod nisi accumsan sit amet. Mauris efficitur egestas mollis. Vestibulum id urna blandit, iaculis tortor sed, hendrerit dui. Proin pulvinar dolor faucibus magna viverra dictum. Donec porttitor pharetra enim, ac blandit sem consectetur nec. Fusce eu metus ac libero posuere tempus id vel dolor. Duis in felis ac enim maximus finibus. Mauris quis ultrices erat, a gravida quam. Nam id mauris nec sem congue tempor."))
			})
		})
	})
}

func New() Interface.MicroAppInterface {
	return &IngexConfigurations{}
}

func (appConfig *IngexConfigurations) index(w http.ResponseWriter, r *http.Request) {
	if(auth.App.Authority.Allow("", r)) {
		w.Write([]byte("Authorised "))
	} else {
		w.Write([]byte("Unauthorised"))
	}
	appConfig.Render.Layout("application18").Execute(path.Join("themes", "theme1", "index"), r.Context(), r, w)
}