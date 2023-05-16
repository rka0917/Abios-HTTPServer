package models

import (
	"fmt"
	"os"

	"github.com/rka0917/Abios-HTTPServer/abios"
)

func GetLiveSeries() ([]map[string]interface{}, error) {

	filters := []string{"lifecycle=live"}
	data, err := abios.GetDataFromEndpoint(fmt.Sprintf("%s/series?filter=", os.Getenv("ABIOSURL")), os.Getenv("ABIOSAPIKEY"), filters)
	if err != nil {
		fmt.Printf("Error occurred sending series request: %s", err)
		return nil, err
	}

	return data, nil
}

func GetSeriesRosters(liveSeriesIds []int) ([]map[string]interface{}, error) {

	var totalData []map[string]interface{}
	for _, liveSeriesId := range liveSeriesIds {
		data, err := abios.GetDataFromEndpoint(fmt.Sprintf("%s/series/%d/rosters?filter=", os.Getenv("ABIOSURL"), liveSeriesId), os.Getenv("ABIOSAPIKEY"), []string{})

		if err != nil {
			fmt.Printf("Error occurred sending series roster request: %s", err)
			return nil, err
		}

		totalData = append(totalData, data...)
	}

	return totalData, nil
}
