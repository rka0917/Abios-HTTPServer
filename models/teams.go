package models

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rka0917/Abios-HTTPServer/abios"
)

func GetTeamsFromIds(teamIds []int) ([]map[string]interface{}, error) {
	var ids []string
	for _, id := range teamIds {
		ids = append(ids, strconv.Itoa(id))
	}

	// Ad-hoc solution. If done properly, only 50 ids at a time should be passed on to the url, to prevent redundant data transfer.
	filters := []string{fmt.Sprintf("id<={%s}", strings.Join(ids, ","))}
	data, err := abios.GetDataFromEndpoint(fmt.Sprintf("%s/teams?filter=", os.Getenv("ABIOSURL")), os.Getenv("ABIOSAPIKEY"), filters)
	if err != nil {
		fmt.Printf("Error occurred sending teams request: %s", err)
		return nil, err
	}

	return data, nil
}
