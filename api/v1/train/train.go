package train

import (
	"github.com/gin-gonic/gin"
	"github.com/junminhong/toot-clock/pkg/collector"
	"github.com/junminhong/toot-clock/pkg/json"
	"log"
	"net/http"
)

var formData = map[string]string{
	"startStation":       "4220-臺南",
	"endStation":         "4210-大橋",
	"transfer":           "ONE",
	"rideDate":           "2021/12/28",
	"startOrEndTime":     "true",
	"startTime":          "00:00",
	"endTime":            "23:59",
	"transferStation":    "0900-基隆",
	"trainTypeList":      "ALL",
	"_isQryEarlyBirdTrn": "on",
}

type searchInfo struct {
	StartStation string `form:"startStation" json:"startStation" binding:"required"`
	EndStation   string `form:"endStation" json:"endStation" binding:"required"`
	RideDate     string `form:"rideDate" json:"rideDate" binding:"required"`
	StartTime    string `form:"startTime" json:"startTime" binding:"required"`
	EndTime      string `form:"endTime" json:"endTime" binding:"required"`
}

func GetTrainInfo(c *gin.Context) {
	query := &searchInfo{}
	err := c.BindJSON(query)
	if err != nil {
		log.Println(err.Error())
	}
	formData["startStation"] = query.StartStation
	formData["endStation"] = query.EndStation
	formData["rideDate"] = query.RideDate
	formData["startTime"] = query.StartTime
	formData["endTime"] = query.EndTime
	trainInfos := collector.TrainInfoCollector(formData)
	c.JSON(http.StatusOK, gin.H{
		"data": trainInfos[0],
	})
}

func GetTrainStation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": json.ProcessStation(),
	})
}
