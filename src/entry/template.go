package entry

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"github.com/Scfy-Code/scfy-im/config"
)

// Views 存储所有的视图
var Views = scan(config.APPCFG.ViewDir)

func scan(dir string) (views map[string]*template.Template) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("扫描%s目录出错！错误信息%s", dir, err.Error())
		os.Exit(2)
	}
	for _, file := range files {
		if file.IsDir() {
			scan(dir + "/" + file.Name())
		} else {
			tmp, err0 := template.ParseFiles(dir + "/" + file.Name())
			if err0 != nil {
				log.Printf("路径为%s的模板解析出现错误！错误信息：%s", dir+""+file.Name(), err0.Error())
			}
			views[file.Name()] = tmp
		}
	}
	return views
}
