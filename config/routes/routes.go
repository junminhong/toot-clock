package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/junminhong/toot-clock/api/v1/train"
	"os"
)

func InitRoutes() {
	router := gin.Default()
	v1 := router.Group("/api/v1/train")
	{
		v1.POST("/time", train.GetTrainInfo)
		v1.GET("/station", train.GetTrainStation)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
