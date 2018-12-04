package handler

import (
	"github.com/gin-gonic/gin"
	"multifort/log"
)

func Status(c *gin.Context) {
	log.Info("[multifort server status info]:", "multifort server is running, everything is ok")
	c.JSON(200, gin.H{
		"code": 0,
		"data": "",
	})
}
