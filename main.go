package main

import (
	"ginDoc/conf"

	_ "ginDoc/db"
	_ "ginDoc/utils"
)

func init() {

}

// @title 还只是个测试
// @description go-gin swagger 文档测试.
// @contact.name zw
// @contact.email wei.zhou@shunwang.com

func main() {
	NewServer().Run(":" + conf.Cfg["httpport"])
}

// doc gen: swagger.exe


//
//package main
//
//import "github.com/gin-gonic/gin"
//
//func main() {
//	gin.DisableConsoleColor()
//	r := gin.Default()
//	r.GET("/hello", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "pong",
//		})
//	})
//	r.Run(":9002") // listen and serve on 0.0.0.0:8080
//}