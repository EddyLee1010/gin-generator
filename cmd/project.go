package cmd

import (
	"fmt"
	"github.com/eddylee1010/gin-generator/generator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"os/exec"
)

var genProjectCmd = &cobra.Command{
	Use:     "project",
	Short:   "创建项目基础目录结构",
	Example: "gin-generator new myapp",
	Long: `Create a new Gin + GORM project with the following directory structure:

	.
	├── cmd
	├── config
	├── controller
	└── router
	`,
	PreRun: func(cmd *cobra.Command, args []string) {
		generator.InitTemplates()
		viper.SetConfigFile("gen-config.yaml")
		err := viper.ReadInConfig()
		if err != nil {
			slog.Error("❌ Failed to read config:", err)
			slog.Error("❌ 请使用gin-generator gen config生成工具所需的配置文件，再次尝试")
			return
		}
		// 检查项目名是否合法
		if !isValidProjectName(viper.GetString("project_name")) {
			slog.Error("Invalid project name. Project name must be a valid Go package name.")
			os.Exit(1)
		}
	},

	Run: func(cmd *cobra.Command, args []string) {
		createProject(viper.GetString("project_name"))
	},
}

// 检查项目名是否合法
func isValidProjectName(s string) bool {
	// 检查是否以字母开头
	if !('a' <= s[0] && s[0] <= 'z' || 'A' <= s[0] && s[0] <= 'Z') {
		return false
	}
	return true
}

func init() {
	genCmd.AddCommand(genProjectCmd)
}

func createProject(name string) {
	// 创建目录结构
	dirs := []string{
		"cmd",
		"config",
		"controller",
		"router",
	}
	for _, dir := range dirs {

		if _, err := os.Stat(dir); err == nil {
			slog.Warn("❌ 项目目录已存在，禁止覆盖！", "path", dir)
			continue
		}
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
		}
	}
	// 渲染输出main.go
	err := generator.RenderTemplateToFile(generator.MainTemplate, nil, "main.go")
	if err != nil {
		fmt.Println("❌ 创建 main.go 失败:", err)
		return
	}
	slog.Info("🤡 Project created successfully!\n")

	cmd := exec.Command("go", "mod", "init", name)
	cmd.Dir = "./" // 设置工作目录为生成的项目目录
	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error("❌ 执行 go mod init 失败", "err", err, "output", string(out))
	} else {
		slog.Info("✅ go mod 创建成功")
	}

	// 执行 go mod tidy todo 将来可自行替换使用位置
	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = "./" // 设置工作目录为生成的项目目录

	out, err = cmd.CombinedOutput()
	if err != nil {
		slog.Error("❌ 执行 go mod tidy 失败", "err", err, "output", string(out))
	} else {
		slog.Info("✅ go mod tidy 成功")
	}

}
