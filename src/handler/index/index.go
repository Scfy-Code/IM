package index

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Scfy-Code/scfy-im/service/index"

	scope "github.com/Scfy-Code/scfy-im/scope/imp"

	"github.com/Scfy-Code/scfy-im/entry"
)

var indexService = index.NewIndexService()

// indexView 首页
type indexView struct {
	redirectURL   string //认证失败的跳转地址
	indexTemplate *template.Template
}

func (iv indexView) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := scope.NewSession(r)
	if !session.IsExist("USER") {
		http.Redirect(w, r, iv.redirectURL, http.StatusFound)
		return
	}
	user := session.GetData("USER")
	friends := indexService.SelectFriends(fmt.Sprintf("%d", user["id"]))
	groups := indexService.SelectGroups(fmt.Sprintf("%d", user["id"]))
	user["friends"] = friends
	user["groups"] = groups
	iv.indexTemplate.Execute(w, user)
}

// NewIndexController 创建一个首页路由器
func NewIndexController() http.Handler {
	return &indexView{"/user/sign_in.scfy", entry.Views["index.scfy"]}
}
