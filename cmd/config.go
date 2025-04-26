package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var genConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "生成cli自身生成所需要的配置文件",
	Run: func(cmd *cobra.Command, args []string) {
		createDefaultConfig()
	},
}

func init() {
	genCmd.AddCommand(genConfigCmd)

}

func createDefaultConfig() {
	const fileName = "gen-config.yaml"

	// 如果文件已存在，退出并提示
	if _, err := os.Stat(fileName); err == nil {
		fmt.Printf("⚠️  配置文件 '%s' 已存在，跳过生成。\n", fileName)
		return
	}

	content := `#      ######   #### ##    ##          ######   ######## ##    ## ######## ########     ###    ########  #######  ######## 
#     ##    ##   ##  ###   ##         ##    ##  ##       ###   ## ##       ##     ##   ## ##      ##    ##     ## ##     ##
#     ##         ##  ####  ##         ##        ##       ####  ## ##       ##     ##  ##   ##     ##    ##     ## ##     ##
#     ##   ####  ##  ## ## ## ####### ##   #### ######   ## ## ## ######   ########  ##     ##    ##    ##     ## ######## 
#     ##    ##   ##  ##  ####         ##    ##  ##       ##  #### ##       ##   ##   #########    ##    ##     ## ##   ##  
#     ##    ##   ##  ##   ###         ##    ##  ##       ##   ### ##       ##    ##  ##     ##    ##    ##     ## ##    ## 
#      ######   #### ##    ##         ######   ######## ##    ## ######## ##     ## ##     ##    ##     #######  ##     ##
#
#	Author: Eddy 527084800

## 🈲🈲🈲🈲 请勿将此文件提交至项目！ 使用 git rm gen-config.yaml --cache 命令从git管理库中移除
## ❗️此文件为cli生成所需配置，config.yaml为生成项目运行配置文件，注意区分

project_name: my-gin-project

database:
  driver: mysql
  dsn: root:123456@tcp(127.0.0.1:3306)/mydb

output:
  model: ./dao
  controller: ./controller
  router: ./router

slog:
  level: error #对应slog的日志级别 error、info、debug
`
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		fmt.Println("❌ 创建配置文件失败:", err)
		return
	}
	fmt.Println("✅ 配置文件 'gin-generator.yaml' 已创建。")
}
