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
