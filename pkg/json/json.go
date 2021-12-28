package json

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type stationInfo struct {
	StationCode   string `json:"stationCode"`
	StationName   string `json:"stationName"`
	StationEName  string `json:"stationEName"`
	Name          string `json:"name"`
	Ename         string `json:"ename"`
	StationAddrTw string `json:"stationAddrTw"`
	StationAddrEn string `json:"stationAddrEn"`
	StationTel    string `json:"stationTel"`
	Gps           string `json:"gps"`
}

func ProcessStation() {
	jsonFile, err := os.Open("station_dataset.json")
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println(err.Error())
	}
	var stationInfos []stationInfo
	err = json.Unmarshal(byteValue, &stationInfos)
	if err != nil {
		panic(err)
	}
	log.Println(stationInfos[0])
}
