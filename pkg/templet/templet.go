package templet

import (
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/Scfy-Code/IM/pkg"
	"github.com/Scfy-Code/IM/pkg/log"
)

var (
	tmp    templet
	logger = log.NewWarnLogger()
)

// templet 模板操作接口
type templet interface {
	// ReturnTemplate 返回指定名称的模板
	returnTemplate(string) *template.Template
}

// ReturnTemplate 返回指定名称的模板
func ReturnTemplate(templateName string) *template.Template {
	return tmp.returnTemplate(templateName)
}

// templetPro 生产环境模板结构体
type templetPro struct {
	templets *template.Template
}

// newTempletPro 创建生产环境的模板对象
func newTempletPro() templet {
	return &templetPro{
		analysisTemplateFiles(analysisTemplateDirs(pkg.APP.TemplateDir)...),
	}
}

func (tp templetPro) returnTemplate(templateName string) *template.Template {
	return tp.templets.Lookup(templateName)
}

// templetDev 开发环境模板结构体
type templetDev struct {
	templateFiles []string
}

// newTempletDev 创建开发环境的模板对象
func newTempletDev() templet {
	return &templetDev{
		analysisTemplateDirs(pkg.APP.TemplateDir),
	}
}
func (td templetDev) returnTemplate(templateName string) *template.Template {
	return analysisTemplateFiles(td.templateFiles...).Lookup(templateName)
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
			logger.Printf("读取模板目录出错！错误信息：%s", err0.Error())
			continue
		}
		for index1 := range files {
			switch files[index1].IsDir() {
			case true:
				templateFiles = append(templateFiles, analysisTemplateDirs(templateDirs[index0]+"/"+files[index1].Name())...)
			case false:
				if strings.Contains(templateDirs[index0]+"/"+files[index1].Name(), ".scfy") {
					templateFiles = append(templateFiles, templateDirs[index0]+"/"+files[index1].Name())
				}
			}
		}
	}
	return templateFiles
}
func init() {
	switch pkg.APP.RuntimeEnv {
	case "PRO":
		tmp = newTempletPro()
	case "DEV":
		tmp = newTempletDev()
	default:
		tmp = newTempletPro()
	}
}
