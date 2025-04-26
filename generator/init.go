package generator

import (
	"github.com/eddylee1010/gin-generator/templates"
	"html/template"
	"log/slog"
	"os"
	"path/filepath"
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
	ServiceTemplate    *template.Template
	DTOTemplate        *template.Template
	MainTemplate       *template.Template // main.go模版
	ConfigTemplate     *template.Template
	ConfigFileTemplate *template.Template
	ControllerTemplate *template.Template
	RouterTemplate     *template.Template
)

// InitTemplates 初始化模板
func InitTemplates() error {
	var err error

	// 初始化service模版
	ServiceTemplate, err = template.New("service").Parse(templates.ServiceTmplStr)
	if err != nil {
		slog.Error("Failed to Parse template", "error", err)
		return err
	}
	// 初始化dto
	DTOTemplate, err = template.New("dto").Parse(templates.DtoTmplStr)
	if err != nil {
		return err
	}

	// 初始化controller模版
	ControllerTemplate, err = template.New("controller").Parse(templates.ControllerTmplStr)
	if err != nil {
		slog.Error("Failed to Parse template", "error", err)
		return err
	}

	// 初始化main模版
	MainTemplate, err = template.New("main").Parse(templates.MainTmplStr)
	if err != nil {
		slog.Error("Failed to Parse template", "error", err)
		return err
	}

	// 初始化config模版
	ConfigTemplate, err = template.New("config").Parse(templates.ConfigTemplateStr)
	if err != nil {
		return err
	}

	// 初始化config.yaml模版
	ConfigFileTemplate, err = template.New("configyaml").Parse(templates.ConfigFileTemplateStr)
	if err != nil {
		return err
	}

	// 初始化router模版
	RouterTemplate, err = template.New("router").Parse(templates.RouterTemplateStr)
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
	os.MkdirAll(filepath.Dir(outPath), os.ModePerm)
	file, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer file.Close()
	slog.Info("📝 生成" + outPath + " success！")
	return tmpl.Execute(file, data)
}
