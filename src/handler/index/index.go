package index

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Scfy-Code/scfy-im/service/index"

	"github.com/Scfy-Code/scfy-im/scope"

	"github.com/Scfy-Code/scfy-im/entry"
)

var indexService = index.NewIndexService()

// indexView 首页
type indexView struct {
	redirectURL   string //认证失败的跳转地址
	indexTemplate *template.Template
}

func (iv indexView) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := scope.NewSession(w, r)
	if !session.IsExist("USER") {
		http.Redirect(w, r, iv.redirectURL, http.StatusFound)
		return
	}
	user := session.GetData("USER")
	if user == nil {
		http.Redirect(w, r, iv.redirectURL, http.StatusFound)
		return
	}
	friends := indexService.SelectFriends(fmt.Sprintf("%s", user["id"]))
	groups := indexService.SelectGroups(fmt.Sprintf("%s", user["id"]))
	indexData := make(map[string]interface{})
	indexData["user"] = user
	indexData["friends"] = friends
	indexData["groups"] = groups
	iv.indexTemplate.Execute(w, indexData)
}

// NewIndexController 创建一个首页路由器
func NewIndexController() http.Handler {
	return &indexView{"/user/sign_in.scfy", entry.Views["index.scfy"]}
}
