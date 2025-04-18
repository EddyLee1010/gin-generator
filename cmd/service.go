package cmd

import (
	"github.com/spf13/cobra"
	"log/slog"
)

var genServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "😘 生成service层代码 😄",
	Run: func(cmd *cobra.Command, args []string) {
		generateServiceFiles()
	},
}

func init() {
	genCmd.AddCommand(genServiceCmd)
}

func generateServiceFiles() {
	slog.Info("🚀 Generating service files [开发中]...")
}
