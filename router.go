package main

import (
	"github.com/gin-gonic/gin"
	"ginDoc/middlewares"
	"ginDoc/handlers/things"
	"ginDoc/conf"
)


func NewServer() *gin.Engine {

	router := gin.Default()
	router.Use(middlewares.Connect)
	router.GET("/things/:_id", things.GetOne)
	router.GET("/things", things.List)
	router.POST("/things", things.Create)
	router.DELETE("/things/:_id", things.Delete)
	router.PUT("/things/:_id", things.Update)

	if conf.Config()["debug"] == "true" {
		router.Static("/doc", "./public/static")
		router.StaticFile("/swagger.json", "./swagger.json")
	}
	return router
}
