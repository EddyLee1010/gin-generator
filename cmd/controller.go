package cmd

import (
	"fmt"
	"github.com/eddylee1010/gin-generator/generator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	genCmd.AddCommand(genServiceCmd)
}

func generateControllerFiles(projectName string) {
	fmt.Println("💘 来太早了。还没开放出来，hhh")
}
