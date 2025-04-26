package templates

const MainTmplStr = `//     ######   #### ##    ##          ######   ######## ##    ## ######## ########     ###    ########  #######  ######## 
//    ##    ##   ##  ###   ##         ##    ##  ##       ###   ## ##       ##     ##   ## ##      ##    ##     ## ##     ##
//    ##         ##  ####  ##         ##        ##       ####  ## ##       ##     ##  ##   ##     ##    ##     ## ##     ##
//    ##   ####  ##  ## ## ## ####### ##   #### ######   ## ## ## ######   ########  ##     ##    ##    ##     ## ######## 
//    ##    ##   ##  ##  ####         ##    ##  ##       ##  #### ##       ##   ##   #########    ##    ##     ## ##   ##  
//    ##    ##   ##  ##   ###         ##    ##  ##       ##   ### ##       ##    ##  ##     ##    ##    ##     ## ##    ## 
//     ######   #### ##    ##         ######   ######## ##    ## ######## ##     ## ##     ##    ##     #######  ##     ##
//
//	Author: Eddy 527084800

package main

import (
	"github.com/gin-gonic/gin"
	"{{.ProjectName}}/config"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"log"
)

func main() {
	if err := config.LoadConfig("config.yaml"); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	dsn:= config.GlobalConfig.Database.User + ":" + config.GlobalConfig.Database.Password + "@tcp(" + config.GlobalConfig.Database.Host+":" + config.GlobalConfig.Database.Port+")/" + config.GlobalConfig.Database.Dbname

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print("❌ 连接数据错误，请正确修改gen-config.yaml中数据库的配置")
		return
	}
	// 全局查询注册 
	dao.Q = dao.Use(db)

	r := gin.Default()

	routers.Init(r)

	r.Run(":" + config.GlobalConfig.Port)
}`
