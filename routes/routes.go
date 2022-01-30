package routes

import (
	"cloudCharts/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {

	router := gin.Default()
	authed := router.Group("", controller.Authentication())
	{
		router.POST("/createDB/", controller.CreateDB)
		router.GET("/getDB/", controller.GetDB)
		authed.GET("/first/", controller.First)
		authed.GET("/second/", controller.Second)
		authed.GET("/third/", controller.Third)
		authed.GET("/forth/", controller.Forth)
	}

	return router
}
