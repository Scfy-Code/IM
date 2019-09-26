package router

import "net/http"

import (
	_ "github.com/Scfy-Code/IM/pkg/conf"
	"github.com/Scfy-Code/IM/pkg/view"
)

type indexTemplate struct {
}

func (*indexTemplate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	view.ReturnTemplate("index.scfy").Execute(w, nil)
}

// NewIndexTemplateRouter 返回首页模板路由
func NewIndexTemplateRouter() http.Handler {
	return &indexTemplate{}
}
