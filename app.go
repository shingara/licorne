package main

import (
	"github.com/gin-gonic/gin"
	"licorne/controllers"
)

func BuildEngine() (*gin.Engine) {
	r := gin.Default()
	r.GET("/v1/games", controllers.GamesIndexHandler)
	r.GET("/v1/games/:id", controllers.GameShowHandler)
	r.POST("/v1/games", controllers.GamesCreateHandler)
	return r
}

func main() {
	BuildEngine().Run(":9000")
}
