package handler

import (
	"log"
	"net/http"
	"text/template"

	"github.com/Scfy-Code/scfy-im/entry"
)

// IndexHandler 首页
type IndexHandler struct {
	ViewPath string
}

func (ih IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//1、验证是否登录(登录后返回页面/未登录重定向登录页面)
	indexEntry := &entry.IndexEntry{}
	//2、
	view, err := template.ParseFiles(ih.ViewPath)
	if err == nil {
		view.Execute(w, indexEntry)
	} else {
		log.Println(err.Error())
	}
}
