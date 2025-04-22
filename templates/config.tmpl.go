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

type Config struct {
	AppName string ` + "`yaml:\"app_name\"`" + `
	Port    int    ` + "`yaml:\"port\"`" + `
	Database struct {
		Host     string ` + "`yaml:\"host\"`" + `
		Port     int    ` + "`yaml:\"port\"`" + `
		User     string ` + "`yaml:\"user\"`" + `
		Password string ` + "`yaml:\"password\"`" + `
		Name     string ` + "`yaml:\"name\"`" + `
	} ` + "`yaml:\"database\"`" + `
}

func LoadConfig(yamlFile string) error{
	viper.SetConfigFile(yamlFile)
	err := viper.ReadInConfig()
	if err != nil {
		slog.Error("❌ Failed to read config:", err)
		return err
	}
	var config Config
	if err = viper.Unmarshal(&config); err != nil {
		slog.Error("❌ 解析配置到结构体错误","key", err.Error())
		return err
	}
	return nil
}
`
