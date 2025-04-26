package cmd

import (
	"fmt"
	"github.com/eddylee1010/gin-generator/generator"
	"github.com/eddylee1010/gin-generator/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var genRouterCmd = &cobra.Command{
	Use:   "router",
	Short: "😘 生成router代码 😄",
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
		generateRouterFiles(viper.GetString("project_name"))
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	genCmd.AddCommand(genRouterCmd)
}

func generateRouterFiles(projectName string) {
	files, err := ioutil.ReadDir("./controllers")
	if err != nil {
		fmt.Println(err)
		return
	}

	var data RouterTmplData
	for _, file := range files {
		// 只要文件，不要文件夹
		if !file.IsDir() {
			// 去掉后缀
			name := file.Name()
			ext := filepath.Ext(name)
			fileName := strings.TrimSuffix(name, ext)
			i := ControllerInfo{
				LowerCamelName: utils.SnakeToLowerCamel(fileName),
				UpperCamelName: utils.TableNameToStructName(fileName),
			}
			data.Controllers = append(data.Controllers, i)
		}
	}
	data.ProjectName = projectName
	generator.RenderTemplateToFile(generator.RouterTemplate, data, "routers/auto.go")

	// 生成自定义的router文件
	customRouterTemplate, _ := template.New("customRouter").Parse(`package routers

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	// 加载自动生成的路由
	InitAutoRouters(r)

	// 如有特殊需求「例如使用中间件」，在下方自定义
	// xxxRouter.Use(func(context *gin.Context) {
	//	do something...
	//	context.Next()
	//}).GET("/test", controller.xxx)
}`)
	generator.RenderTemplateToFile(customRouterTemplate, nil, "routers/custom.go")
	//}
}

type ControllerInfo struct {
	LowerCamelName string // 小驼峰名
	UpperCamelName string // 大驼峰名
}

type RouterTmplData struct {
	ProjectName string
	Controllers []ControllerInfo
}
