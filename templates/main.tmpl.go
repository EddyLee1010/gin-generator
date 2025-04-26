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
	"log"
)

func main() {
	if err := config.LoadConfig("config.yaml"); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	// 全局查询注册 
	dao.Q = dao.Use(db)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from gin-generator!"})
	})
	r.Run()
}`
