package main

import (
	"github.com/gin-gonic/gin"
	boot "be_patrol/server/bootstrap"
)


func main() {

	r := gin.Default()
	bootApp := boot.Bootup{
		R : r,
	}

	bootApp.RegisterRoutes()

	r.Run("127.0.0.1:2121")
	
}