package templates

// ConfigTemplateStr  生成config.yaml对应的结构体文件的模板
const ConfigTemplateStr = `//     ######   #### ##    ##          ######   ######## ##    ## ######## ########     ###    ########  #######  ######## 
//    ##    ##   ##  ###   ##         ##    ##  ##       ###   ## ##       ##     ##   ## ##      ##    ##     ## ##     ##
//    ##         ##  ####  ##         ##        ##       ####  ## ##       ##     ##  ##   ##     ##    ##     ## ##     ##
//    ##   ####  ##  ## ## ## ####### ##   #### ######   ## ## ## ######   ########  ##     ##    ##    ##     ## ######## 
//    ##    ##   ##  ##  ####         ##    ##  ##       ##  #### ##       ##   ##   #########    ##    ##     ## ##   ##  
//    ##    ##   ##  ##   ###         ##    ##  ##       ##   ### ##       ##    ##  ##     ##    ##    ##     ## ##    ## 
//     ######   #### ##    ##         ######   ######## ##    ## ######## ##     ## ##     ##    ##     #######  ##     ##
//
//	Author: Eddy 527084800

package config

import (
	"github.com/spf13/viper"
	"log/slog"
)
var GlobalConfig Config
type Config struct {
	AppName string ` + "`yaml:\"app_name\"`" + `
	Port    string    ` + "`yaml:\"port\"`" + `
	Database struct {
		Host     string ` + "`yaml:\"host\"`" + `
		Port     string    ` + "`yaml:\"port\"`" + `
		User     string ` + "`yaml:\"user\"`" + `
		Password string ` + "`yaml:\"password\"`" + `
		Dbname     string ` + "`yaml:\"dbname\"`" + `
	} ` + "`yaml:\"database\"`" + `
}

func LoadConfig(yamlFile string) error{
	viper.SetConfigFile(yamlFile)
	err := viper.ReadInConfig()
	if err != nil {
		slog.Error("❌ Failed to read config:", err)
		return err
	}
	if err = viper.Unmarshal(&GlobalConfig); err != nil {
		slog.Error("❌ 解析配置到结构体错误","key", err.Error())
		return err
	}
	return nil
}
`
