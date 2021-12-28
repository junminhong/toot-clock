package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/junminhong/toot-clock/api/v1/train"
)

func InitRoutes() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/train", train.GetTrainInfo)
	}
	router.Run(":8080")
}
