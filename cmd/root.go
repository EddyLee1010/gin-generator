package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gin-generator",
	Short: "🕹️一款自动创建基于gin+GORM项目的cli小工具😂",
	Long: `💡 gin-generator🐔 可以快速把你建立起一个gin+GORM的项目
💡 包括自动生成项目结构、数据库模型+query、service+DTO、controller、router的代码`,
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "生成命令 🔑help 获取使用方法",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "获取当前版本号",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.2.0")
	},
}

func init() {
	rootCmd.AddCommand(genCmd) // 添加gen子命令
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
