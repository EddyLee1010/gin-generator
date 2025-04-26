package cmd

import (
	"github.com/eddylee1010/gin-generator/generator"
	"github.com/eddylee1010/gin-generator/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"path/filepath"
	"strings"
	"unicode"
)

var genControllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "😘 生成controller层代码 😄",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := generator.InitTemplates()

		viper.SetConfigFile("gen-config.yaml")
		err = viper.ReadInConfig()
		if err != nil {
			return err
		}

		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		generateControllerFiles(viper.GetString("project_name"))
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	genCmd.AddCommand(genControllerCmd)
}

func generateControllerFiles(projectName string) {
	files, err := ioutil.ReadDir("./services")
	if err != nil {
		return
	}

	for _, file := range files {
		// 只要文件，不要文件夹
		if !file.IsDir() {
			// 去掉后缀
			name := file.Name()
			ext := filepath.Ext(name)
			fileName := strings.TrimSuffix(name, ext)
			data := ExtractControllerTemplateDataFromService(fileName, "comment名称", projectName)
			generator.RenderTemplateToFile(generator.ControllerTemplate, data, "controllers/"+fileName+".go")
		}
	}

	//}
}

type ControllerTemplateData struct {
	ProjectName string
	StructName  string // 如 BasicFileCat
	VarName     string // 如 basicFileCat
	Comment     string // 例如 "文件分类"
}

func UpperCamelToLowerCamel(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

// 基于 Service 名字生成 Controller 模板数据
func ExtractControllerTemplateDataFromService(structName string, comment string, projectName string) *ControllerTemplateData {
	return &ControllerTemplateData{
		ProjectName: projectName,
		StructName:  utils.SnakeToUpperCamel(utils.TableNameToStructName(structName)),
		VarName:     utils.SnakeToUpperCamel(utils.TableNameToStructName(structName)),
		Comment:     comment,
	}
}
