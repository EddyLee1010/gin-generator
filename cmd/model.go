package cmd

import (
	"fmt"
	"github.com/eddylee1010/gin-generator/generator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"html/template"
	"log/slog"
	"os"
)

var dsn string

var genModelCmd = &cobra.Command{
	Use:   "model",
	Short: "😘 Generate GORM models from an existing database 😄",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if _, err := os.Stat("gen-config.yaml"); os.IsNotExist(err) {
			slog.Error("❌ 请使用gin-generator gen config生成工具所需的配置文件，再次尝试")
			return err
		}
		viper.SetConfigFile("gen-config.yaml")
		err := viper.ReadInConfig()
		if err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		generateModelsFromConfig()
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	//genModelCmd.Flags().StringVar(&dsn, "dsn", "", "MySQL DSN, e.g. user:pass@tcp(127.0.0.1:3306)/dbname")
	genCmd.AddCommand(genModelCmd)
}

func generateModelsFromConfig() {

	dsn = viper.GetString("database.dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print("❌ 连接数据错误，请正确修改gen-config.yaml中数据库的配置")
		return
	}

	outPath := viper.GetString("output.model")
	// 如果目录不存在，先创建
	if err := os.MkdirAll(outPath, os.ModePerm); err != nil {
		slog.Error("❌ 无法创建目录 [%s]：%v\n", outPath, err)
		return
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:        outPath,
		ModelPkgPath:   outPath + "/models",
		Mode:           gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:  true,
		FieldCoverable: true,
	})
	g.UseDB(db)

	slog.Info("🚀 Generating models...")
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
	slog.Info("✅ Models generated.")

	// 初始化项目的db,生成Init.go
	InitDbTemplate, err := template.New("initDb").Parse(`package dao

func InitDb() {

	db, err := gorm.Open(mysql.Open(config.GlobalConfig.Database.User+":"+config.GlobalConfig.Database.Password+"@tcp("+config.GlobalConfig.Database.Host+":"+fmt.Sprintf("%d", config.GlobalConfig.Database.Port)+")/"+config.GlobalConfig.Database.Name+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Printf("❌ 数据库连接失败: %v", err)
	}
	Use(db) // 初始化全局 Query 变量
}`)
	if err != nil {
		fmt.Println(err)
	}
	err = generator.RenderTemplateToFile(InitDbTemplate, nil, "dao/initDb.go")
	if err != nil {
		return
	}
}
