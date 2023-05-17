package abios

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var xRemainingRate = 5 // Chose an arbitrary number > 0 as initial value

type AbiosClient struct {
	RequestURL string
	HTTPClient *http.Client
	APIKey     string
}

func CreateAbiosClient(url string, apiKey string) *AbiosClient {
	return &AbiosClient{
		RequestURL: url,
		HTTPClient: &http.Client{},
		APIKey:     apiKey,
	}
}

// Sends GET requests with regards to paginated anbswers.
func GetDataFromEndpoint(endpoint string, apiKey string, filters []string) ([]map[string]interface{}, error) {
	pageSize := 50
	skip := 0
	var lastPage []map[string]interface{}
	var totalData []map[string]interface{}

	// Need to repeat the GET request for pagination purposes
	for {
		// For now, we use a fixed delay for when we surpass the rate limit
		if xRemainingRate <= 0 {
			time.Sleep(1 * time.Second)
		}

		filterString := strings.Join(filters[:], "&") + fmt.Sprintf("&skip=%d", skip)
		requestURL := fmt.Sprintf("%s%s", endpoint, filterString)
		client := CreateAbiosClient(requestURL, apiKey)
		data, rateRemain, err := client.SendGetRequest()
		xRemainingRate = rateRemain
		lastPage = data

		if err != nil {
			fmt.Printf("Error occurred sending request: %s", err)
			return nil, err
		}

		if len(data) == 0 {
			break
		}

		totalData = append(totalData, data...)

		// When we know we have received the last page
		if len(lastPage) < pageSize {
			break
		}

		skip += pageSize

	}

	return totalData, nil
}

// Sends GET request to endpoint specified in the the AbiosClient
func (c *AbiosClient) SendGetRequest() ([]map[string]interface{}, int, error) {
	req, err := http.NewRequest("GET", c.RequestURL, nil)
	if err != nil {
		fmt.Printf("Error when creating request object: %s\n", err.Error())
		return nil, 0, err
	}

	var data []map[string]interface{}
	req.Header.Add("Abios-Secret", c.APIKey)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Printf("Error when sending HTTP GET Request: %s\n", err.Error())
		return nil, 0, err
	}
	rateRemaining, _ := strconv.Atoi(resp.Header.Get("X-RateLimit-Remaining"))

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error when reading response: %s\n", err.Error())
		return nil, 0, err
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Printf("Error when parsing json response: %s\n", err.Error())
		return nil, 0, err
	}
	return data, rateRemaining, err
}
