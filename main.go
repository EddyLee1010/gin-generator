package main

import (
	"github.com/eddylee1010/gin-generator/cmd"
)

func main() {

	//data := map[string]any{
	//	"PackageName": "front", // 或 dynamic，根据配置
	//	"StructName":  "CmsBanner",
	//	"ProjectPath": "go-cms",
	//	//"FilterFields": [...],               // 从 model 中提取
	//	//"ExportedFields": [...],             // 字段列表
	//	//"ExportedFieldsNoID": [...],         // 除 ID 外字段
	//}
	//
	//err := generator.InitTemplates()
	//if err != nil {
	//	log.Fatalf("failed to render template: %v", err)
	//	return
	//}
	//err = generator.RenderTemplateToFile(generator.RequestServiceTemplate, data, "services/front/banner.go")
	//if err != nil {
	//	log.Fatalf("failed to render template: %v", err)
	//}

	cmd.Execute()
}
