package sys

import (
	"html/template"
	"io/ioutil"
	"strings"
)

// newTemplate 创建一个视图解析器
func newTemplate() Template {
	switch APP.RuntimeEnv {
	case "PRO":
		return &templatePRO{
			analysisTemplateFiles(
				analysisTemplateDirs(
					APP.TemplateDir,
				)...,
			),
		}
	case "DEV":
		return &templateDEV{
			analysisTemplateDirs(
				APP.TemplateDir,
			),
		}
	default:
		return &templatePRO{
			analysisTemplateFiles(
				analysisTemplateDirs(
					APP.TemplateDir,
				)...,
			),
		}
	}
}
func init() {
	temp = newTemplate()
}

// ReturnTemplate 返回与名称对应的模板
func ReturnTemplate(templateName string) *template.Template {
	return temp.ReturnTemplate(templateName)
}

// Template 视图接口
type Template interface {
	// 返回与名称对应的模板
	ReturnTemplate(templateName string) *template.Template
}

// templatePRO 生产环境视图模板
type templatePRO struct {
	templates *template.Template
}

func (temp templatePRO) ReturnTemplate(templateName string) *template.Template {
	return temp.templates.Lookup(templateName)
}

// templateDEV 开发环境视图模板
type templateDEV struct {
	templateFiles []string
}

func (temp templateDEV) ReturnTemplate(templateName string) *template.Template {
	return analysisTemplateFiles(temp.templateFiles...).Lookup(templateName)
}

// analysisTemplateFiles 解析模板
func analysisTemplateFiles(templateFiles ...string) *template.Template {
	templates, err := template.ParseFiles(templateFiles...)
	if err != nil {
		return nil
	}
	return templates
}

// analysisTemplateDirs 解析模板目录获取所有的模板地址
func analysisTemplateDirs(templateDirs ...string) []string {
	var templateFiles []string
	for index0 := range templateDirs {
		files, err0 := ioutil.ReadDir(templateDirs[index0])
		if err0 != nil {
			WarnLogger.Printf("读取模板目录出错！错误信息：%s", err0.Error())
			continue
		}
		for index1 := range files {
			switch files[index1].IsDir() {
			case true:
				templateFiles = append(templateFiles, analysisTemplateDirs(templateDirs[index0]+"/"+files[index1].Name())...)
			case false:
				if strings.Contains(files[index1].Name(), ".scfy") {
					templateFiles = append(templateFiles, templateDirs[index0]+"/"+files[index1].Name())
				}
				continue
			}
		}
	}
	return templateFiles
}
