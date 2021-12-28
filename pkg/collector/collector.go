package collector

import (
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
)

var trainInfos []trainInfo

type trainInfo struct {
	Index         string
	TrainName     string
	StartStation  string
	EndStation    string
	DepartureTime string
	ArrivalTime   string
	DrivingTime   string
	Route         string
	FullTix       string
	ChildTix      string
	OlderTix      string
}

func trainInfoHandler(infos [][]string) []trainInfo {
	trainInfos = []trainInfo{}
	for i := 0; i < len(infos); i++ {
		trainInfos = append(trainInfos, trainInfo{
			Index:         infos[i][0],
			TrainName:     infos[i][1],
			StartStation:  infos[i][3],
			EndStation:    infos[i][5],
			DepartureTime: infos[i][7],
			ArrivalTime:   infos[i][8],
			DrivingTime:   infos[i][9],
			Route:         infos[i][11],
			FullTix:       infos[i][14],
			ChildTix:      infos[i][16],
			OlderTix:      infos[i][18],
		})
	}
	return trainInfos
}

func TrainInfoCollector(formData map[string]string) []trainInfo {
	c := colly.NewCollector()
	url := "https://www.railway.gov.tw/tra-tip-web/tip/tip001/tip112/querybytimeblank"
	c.Visit(url)
	var infos [][]string
	c.OnHTML(".trip-column", func(e *colly.HTMLElement) {
		infos = append(infos, strings.Fields(e.Text))
	})
	err := c.Post(url, formData)
	if err != nil {
		log.Println(err)
	}
	return trainInfoHandler(infos)
}
