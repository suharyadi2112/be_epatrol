package bootstrap

import (
	"github.com/gin-gonic/gin"
)


func (boot *Bootup) RegisterRoutes() {

	v1 := boot.R.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

}