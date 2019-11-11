package router

import (
	"fmt"
	"net/http"

	//只使用初始化方法
	_ "github.com/Scfy-Code/IM/pkg/conf"
	"github.com/Scfy-Code/IM/pkg/view"
)

type indexTemplate struct {
}

func (indexTemplate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		data       map[string][]map[string]interface{} = make(map[string][]map[string]interface{})
		talkerList []map[string]interface{}            = make([]map[string]interface{}, 20)
		groupList  []map[string]interface{}            = make([]map[string]interface{}, 20)
	)
	for index, talker := range talkerList {
		talker = make(map[string]interface{})
		talker["talkerID"] = fmt.Sprintf("%b", index)
		talker["talkerAvatar"] = "/static/images/talker.png"
		talker["talkerNickName"] = "李二狗"
		talker["talkerSign"] = "智乱天下，武逆乾坤！"
		talker["status"] = true
		talker["msgContent"] = "你好啊"
		talkerList[index] = talker
	}
	for index, group := range groupList {
		group = make(map[string]interface{})
		group["groupID"] = fmt.Sprintf("%b", index)
		group["groupAvatar"] = "/static/images/talker.png"
		group["groupNickName"] = "李二狗"
		group["groupSign"] = "智乱天下，武逆乾坤！"
		group["status"] = true
		group["msgContent"] = "你好啊"
		groupList[index] = group
	}
	data["groupList"] = groupList
	data["talkerList"] = talkerList
	view.ReturnTemplate("index.scfy").Execute(w, data)
}

// NewIndexTemplateRouter 返回首页模板路由
func NewIndexTemplateRouter() http.Handler {
	return &indexTemplate{}
}
