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

func GetDataFromEndpoint(endpoint string, apiKey string, filters []string) ([]interface{}, error) {
	pageSize := 50
	skip := 0
	var lastPage []interface{}
	var totalData []interface{}
	filterString := strings.Join(filters[:], "&")
	for {
		requestURL := fmt.Sprintf("%s%s&skip=%d", endpoint, filterString, skip)
		client := CreateAbiosClient(requestURL, apiKey)
		data, rateRemain, err := client.SendGetRequest()
		lastPage = data

		if err != nil {
			fmt.Printf("Error occurred sending request: %s", err)
			return nil, err
		}

		if len(data) == 0 {
			break
		}

		totalData = append(totalData, data...)

		if len(lastPage) < pageSize {
			break
		}

		skip += pageSize

		// For now, we use a fixed delay for when we surpass the rate limit
		if rateRemain <= 0 {
			time.Sleep(1 * time.Second)
		}
	}

	return totalData, nil
}

func (c *AbiosClient) SendGetRequest() ([]interface{}, int, error) {
	req, err := http.NewRequest("GET", c.RequestURL, nil)
	if err != nil {
		fmt.Printf("Error when creating request object: %s\n", err.Error())
		return nil, 0, err
	}

	var data []interface{}
	req.Header.Add("Abios-Secret", c.APIKey)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Printf("Error when sending HTTP GET Request: %s\n", err.Error())
		return nil, 0, err
	}
	rateRemaining, _ := strconv.Atoi(resp.Header.Get("X-Ratelimit-Remaining"))

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
