package server

import (
	"github.com/gin-gonic/gin"
	"github.com/denislee/stack-backend/controller"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	ping := new(controller.PingController)
	router.GET("/ping", ping.Status)

	docker := new(controller.DockerController)
	router.GET("/docker/status", docker.Status)
	router.GET("/docker/containers", docker.Containers)

	return router
}
