package utils

import (
	"github.com/qor/render"
	"net/http"
	"html/template"
)

func CopyRender(r *render.Render, funcmaps... interface{}) *render.Render {
	vp := r.ViewPaths
	vp = append(vp, "app/home/views")
	Render := render.New(r.Config, vp...)
	oldFuncMapMaker := r.FuncMapMaker
	Render.FuncMapMaker = func(r *render.Render, req *http.Request, w http.ResponseWriter) template.FuncMap {
		funcMap := template.FuncMap{}
		if oldFuncMapMaker != nil {
			for k, v := range oldFuncMapMaker(r, req, w) {
				funcMap[k] = v
			}
		}

		for i := 0; i < len(funcmaps); i++ {
			switch funcmaps[i].(type) {
			case map[string]interface{}:
				fm := funcmaps[i].(map[string]interface{})
				for k, v := range fm {
					funcMap[k] = v
				}
				break
			case func(*render.Render, *http.Request, http.ResponseWriter) template.FuncMap:
				fn := funcmaps[i].(func(*render.Render, *http.Request, http.ResponseWriter) template.FuncMap)
				for k, v := range fn(r, req, w) {
					funcMap[k] = v
				}
			case string:
				Render.Config.DefaultLayout = funcmaps[i].(string)
				break
			default:
				panic("funcmaps is nor map[string]interface{}")
				break
			}

		}

		return funcMap
	}
	return Render
}

