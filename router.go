package main

import (
	"github.com/gin-gonic/gin"
	"ginDoc/middlewares"
	"ginDoc/handlers/sale"

	"ginDoc/conf"
)


func NewServer() *gin.Engine {


	router := gin.Default()
	router.Use(middlewares.Connect)
	sale.ThingContoller{}.SetRoute(router)
	sale.ArticleController{}.SetRoute(router)

	if conf.Config()["debug"] == "true" {
		router.Static("/doc", "./public/static")
		router.StaticFile("/swagger.json", "./swagger.json")
	}
	return router
}
