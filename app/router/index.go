package router

import (
	"net/http"

	_ "github.com/Scfy-Code/IM/pkg/conf"
	"github.com/Scfy-Code/IM/pkg/view"
)

type indexTemplate struct {
}

func (*indexTemplate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var userList []map[string]interface{} = make([]map[string]interface{}, 10)
	for _, user := range userList {
		user = make(map[string]interface{})
		user["userAvatar"] = "/static/images/avatar.png"
		user["userNickName"] = "李二狗"
		user["userSign"] = "智乱天下，武逆乾坤！"
	}
	view.ReturnTemplate("index.scfy").Execute(w, userList)
}

// NewIndexTemplateRouter 返回首页模板路由
func NewIndexTemplateRouter() http.Handler {
	return &indexTemplate{}
}
