package models

import (
	"fmt"
	"os"

	"github.com/rka0917/Abios-HTTPServer/abios"
)

func GetLiveSeries() ([]interface{}, error) {

	filters := []string{"lifecycle=live"}
	data, err := abios.GetDataFromEndpoint(fmt.Sprintf("%s/series?filter=", os.Getenv("ABIOSURL")), os.Getenv("ABIOSAPIKEY"), filters)
	if err != nil {
		fmt.Printf("Error occurred sending request: %s", err)
		return nil, err
	}

	return data, nil
}
