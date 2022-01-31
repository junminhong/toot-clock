package train

import (
	"github.com/gin-gonic/gin"
	"github.com/junminhong/toot-clock/pkg/cleaner"
	"github.com/junminhong/toot-clock/pkg/collector"
	"log"
	"net/http"
)

var formData = map[string]string{
	"startStation": "4220-臺南",
	"endStation":   "4210-大橋",
	// ONE NORMAL
	"transfer":        "ONE",
	"rideDate":        "2021/12/28",
	"startOrEndTime":  "true",
	"startTime":       "00:00",
	"endTime":         "23:59",
	"transferStation": "0900-基隆",
	//ALL=全部 RESERVED_TRAIN=對號 NON_RESERVED=非對號
	"trainTypeList":      "ALL",
	"_isQryEarlyBirdTrn": "on",
}

type searchInfo struct {
	StartStation  string `json:"start_station" binding:"required"`
	EndStation    string `json:"end_station" binding:"required"`
	RideDate      string `json:"ride_date" binding:"required"`
	StartTime     string `json:"start_time" binding:"required"`
	EndTime       string `json:"end_time" binding:"required"`
	TrainTypeList string `json:"train_type_list" binding:"required"`
}

type station struct {
	Station string `json:"Station"`
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
	formData["trainTypeList"] = query.TrainTypeList
	trainInfos := collector.TrainInfoCollector(formData)
	c.JSON(http.StatusOK, gin.H{
		"data": trainInfos,
	})
}

func GetTrainStation(c *gin.Context) {
	stationInfos := cleaner.ProcessStation()
	var stations []station
	for index := 0; index < len(stationInfos); index++ {
		stations = append(stations, station{Station: stationInfos[index].StationCode + stationInfos[index].StationName})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": stations,
	})
}
