package app

import (
	"html/template"
	"io/ioutil"
	"os"
)

var (
	// TemplateMap 模板字典
	TemplateMap = make(map[string]*template.Template)
)

func newTemplateMap(templateDir string) {
	files, err := ioutil.ReadDir(templateDir)
	if err != nil {
		ErrorLogger.Printf("扫描%s目录出错！错误信息%s", templateDir, err.Error())
		os.Exit(2)
	}
	for _, file := range files {
		if file.IsDir() {
			newTemplateMap(templateDir + "/" + file.Name())
		} else {
			tmp, err0 := template.ParseFiles(templateDir + "/" + file.Name())
			if err0 != nil {
				ErrorLogger.Printf("路径为%s的模板解析出现错误！错误信息：%s", templateDir+"/"+file.Name(), err0.Error())
			}
			TemplateMap[file.Name()] = tmp
		}
	}
}
