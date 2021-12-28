package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/junminhong/toot-clock/api/v1/train"
)

func InitRoutes() {
	router := gin.Default()
	v1 := router.Group("/api/v1/train")
	{
		v1.POST("/time", train.GetTrainInfo)
		v1.GET("/station", train.GetTrainStation)
	}
	router.Run(":8080")
}
