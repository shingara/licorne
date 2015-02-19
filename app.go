package main

import (
	"github.com/gin-gonic/gin"
	"licorne/controllers"
)


func main() {
	r := gin.Default()
	r.GET("/v1/games", controllers.GamesIndexHandler)
	r.POST("/v1/games", controllers.GamesCreateHandler )
	r.Run(":9000")
}
