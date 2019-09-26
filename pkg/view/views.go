package view

import (
	"html/template"
	"io/ioutil"
	"os"

	"github.com/Scfy-Code/IM/pkg/log"
)

var templates *template.Template
var logger = log.NewWarnLogger()

func scanTemplateDir(templateDir string) (templateList []string) {
	files, err := ioutil.ReadDir(templateDir)
	if err != nil {
		logger.Println(err.Error())
		os.Exit(2)
	}
	for _, file := range files {
		switch file.IsDir() {
		case true:
			templateList = append(templateList, scanTemplateDir(templateDir+"/"+file.Name())...)
		case false:
			templateList = append(templateList, templateDir+"/"+file.Name())
		}
	}
	return
}
func createTemplates(templateList []string) *template.Template {
	t, e := template.ParseFiles(templateList...)
	if e != nil {
		logger.Println(e.Error())
		os.Exit(2)
	}
	return t
}

// RegistTemplateDir 注册一个模板扫描路径
func RegistTemplateDir(templateDir string) {
	templates = createTemplates(scanTemplateDir(templateDir))
}

// ReturnTemplate 给客户端返回指定名称的页面
func ReturnTemplate(templateName string) *template.Template {
	return templates.Lookup(templateName)
}
