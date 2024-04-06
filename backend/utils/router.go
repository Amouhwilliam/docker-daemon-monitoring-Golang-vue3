package utils

import (
	"github.com/gin-gonic/gin"
	"kinexon/containerruntime/app/controllers"
	"kinexon/containerruntime/middlewares"
)

func GetRouter() *gin.Engine {
	router := gin.New()
	router.Use(middlewares.CORS())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/info", controllers.Info)
	dockerRouter := router.Group("/containers")

	dockerRouter.GET("/json", controllers.ListContainer)
	dockerRouter.POST("/:id/restart", controllers.RestartContainer)
	dockerRouter.DELETE("/:id", controllers.RemoveContainer)
	dockerRouter.POST("/:id/stop", controllers.StopContainer)
	dockerRouter.GET("/:id/stats", controllers.StartStreamContainerStats)

	return router
}
