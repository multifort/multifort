package router

import (
	"github.com/gin-gonic/gin"
	"multifort/pkg/handler"
)

func Init(router *gin.Engine)  {
	router.Use(
		gin.Recovery(),
	)

	api := router.Group("/api")
	{
		api.GET("/status",handler.Status)
	}

	api = api.Group("/v1")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
	}

}