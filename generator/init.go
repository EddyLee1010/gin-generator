package generator

import (
	"html/template"
	"log/slog"
	"os"
	"path/filepath"
)

var (
	RequestServiceTemplate *template.Template
	DTOTemplate            *template.Template
	MainTemplate           *template.Template // main.go模版
)

// InitTemplates 初始化模板
func InitTemplates() error {
	var err error
	RequestServiceTemplate, err = LoadTemplate("templates/service.tmpl")
	if err != nil {
		return err
	}
	DTOTemplate, err = LoadTemplate("templates/dto.tmpl")
	if err != nil {
		return err
	}
	MainTemplate, err = LoadTemplate("templates/main.go.tmpl")
	if err != nil {
		return err
	}
	return nil
}

func LoadTemplate(path string) (*template.Template, error) {
	tmpl, err := template.New(filepath.Base(path)).ParseFiles(path)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
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
