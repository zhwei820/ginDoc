package main

import (
	"github.com/chrisjstevenson/go-gin-mgo/db"
	"github.com/chrisjstevenson/go-gin-mgo/conf"
)


func init() {
	db.Connect()

}

// @title 还只是个测试
// @description go-gin swagger 文档测试.
// @contact.name zw
// @contact.email wei.zhou@shunwang.com

func main() {
	NewServer().Run(":" + conf.Config()["httpport"])
}