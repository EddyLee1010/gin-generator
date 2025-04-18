package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gin-generator",
	Short: "🕹️一款自动创建基于gin+GORM项目的cli小工具😂",
	Long: `💡 gin-generator🐔 可以快速把你建立起一个gin+GORM的项目
💡 包括自动生成数据库模型、service+DTO、controller、router的代码`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
