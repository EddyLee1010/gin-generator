package generator

import (
	"github.com/eddylee1010/gin-generator/templates"
	"html/template"
	"log/slog"
	"os"
)

// TemplateConfigData 模板配置结构体，cli使用。不是用户项目数据
type TemplateConfigData struct {
	ProjectName string
	Port        int
	Database    struct {
		DBHost     string
		DBPort     string
		DBUser     string
		DBPassword string
		DBName     string
	}
}

var (
	RequestServiceTemplate *template.Template
	DTOTemplate            *template.Template
	MainTemplate           *template.Template // main.go模版
	ConfigTemplate         *template.Template
	ConfigFileTemplate     *template.Template
)

// InitTemplates 初始化模板
func InitTemplates() error {
	var err error
	//RequestServiceTemplate, err = LoadTemplate("templates/service.tmpl")
	//if err != nil {
	//	return err
	//}
	//DTOTemplate, err = LoadTemplate("templates/dto.tmpl")
	//if err != nil {
	//	return err
	//}
	MainTemplate, err = template.New("main").Parse(templates.MainTmplStr)
	if err != nil {
		slog.Error("Failed to Parse template", "error", err)
		return err
	}
	ConfigTemplate, err = template.New("config").Parse(templates.ConfigTemplateStr)
	if err != nil {
		return err
	}
	ConfigFileTemplate, err = template.New("config.yaml").Parse(templates.ConfigFileTemplateStr)
	if err != nil {
		return err
	}
	return nil
}

// RenderTemplateToFile 将模板渲染并写入文件
// tmpl: 模板对象
// data: 模板数据
// outPath: 输出文件路径
func RenderTemplateToFile(tmpl *template.Template, data any, outPath string) error {
	file, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer file.Close()
	slog.Info("📝 生成" + outPath + " success！")
	return tmpl.Execute(file, data)
}
