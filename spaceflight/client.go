package spaceflight

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// https://api.spaceflightnewsapi.net/v3
/*
{
	"/articles/count",
	"/articles",
	"/articles/{id}",
	"/articles/launch/{id}",
	"/articles/event/{id}",
	"/blog/count",
	"/blogs",
	"/blogs/{id}",
	"/blogs/launch/{id}",
	"/blogs/event/{id}"
}
*/

// _limit=10
// _start=
// _contains

// 'accept: application/json'

type UrlParams map[string]string

func (params UrlParams) Query() string {
	query := ""

	for key, value := range params {
		if len(query) > 0 {
			query += "&"
		}
		query += key + "=" + value
	}
	/*
		args := make([]string, 10)
		for key, value := range params {
			args = append(args, key+"="+value)
		}
		query = strings.Join(args, "&")
	*/
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
			Timeout: time.Second,
		},
	}
}

func PrintResponse(res *http.Response) {
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Printf("%v\n", bodyString)
}

func (client *Client) SendRequest(request *http.Request, result interface{}) error {
	if len(client.ApiKey) > 0 {
		request.Header.Add("X-Api-Key", client.ApiKey)
	}

	res, err := client.HTTPClient.Do(request)
	if err != nil {
		return fmt.Errorf("client: %s", err.Error())
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request, code: %s", res.Status)
	} else {
		fmt.Println(res.StatusCode)
		//PrintResponse(res)
	}

	err = json.NewDecoder(res.Body).Decode(result)
	if err != nil {
		return fmt.Errorf("decode %s", err.Error())
	}
	return nil
}
