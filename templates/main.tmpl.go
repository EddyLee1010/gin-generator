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
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from gin-generator!"})
	})
	r.Run()
}`
