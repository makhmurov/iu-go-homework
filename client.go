package finder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type UrlParams map[string]string

func (params UrlParams) Query() string {
	query := ""

	for key, value := range params {
		if len(query) > 0 {
			query += "&"
		}
		query += key + "=" + value
	}
	return query
}

type Client struct {
	BaseUrl    string
	ApiKey     string
	HTTPClient *http.Client
}

func NewClient(url, key string) *Client {
	return &Client{
		BaseUrl: url,
		ApiKey:  key,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (client *Client) SendRequest(request *http.Request, result interface{}) error {
	if len(client.ApiKey) > 0 {
		request.Header.Add("X-Api-Key", client.ApiKey)
	}

	res, err := client.HTTPClient.Do(request)
	if err != nil {
		return fmt.Errorf("client: %s", err.Error())

	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request, code: %s", res.Status)
	} else {
		fmt.Println(res.StatusCode)
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(result)
	if err != nil {
		return fmt.Errorf("decode %s", err.Error())
	}
	return nil
}
