package entry

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

// Views 存储所有的视图
var Views = make(map[string]*template.Template)

func scan(dir string) {
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
			Views[file.Name()] = tmp
		}
	}
}
func init() {
	scan("../lib/views")
}
