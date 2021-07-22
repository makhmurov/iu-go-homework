package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

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

type NewsInfo struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source struct {
			ID   interface{} `json:"id"`
			Name string      `json:"name"`
		} `json:"source"`
		Author      interface{} `json:"author"`
		Title       string      `json:"title"`
		Description string      `json:"description"`
		URL         string      `json:"url"`
		URLToImage  string      `json:"urlToImage"`
		PublishedAt time.Time   `json:"publishedAt"`
		Content     string      `json:"content"`
	} `json:"articles"`
}

func main() {
	fmt.Println("hi")
	client := NewClient("https://newsapi.org/v2", os.Getenv("DATA_PROVIDER_APIKEY"))

	var params map[string]string = map[string]string{
		"q":        "tesla",
		"form":     "2021-06-21",
		"sortBy":   "publishedAt",
		"apiKey":   client.ApiKey,
		"pageSize": "5",
		"language": "ru",
	}

	query := ""

	for key, value := range params {
		if len(query) > 0 {
			query += "&"
		}
		query += key + "=" + value
	}

	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/everything?%s", client.BaseUrl, query),
		nil,
	)

	_ = err

	res, err := client.HTTPClient.Do(request)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}

	fmt.Println(res.StatusCode, ":", res.Status)

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error request:", res.StatusCode)
		return
	}

	// var buf []byte = make([]byte, 100)
	// count, err := res.Body.Read(buf)
	// fmt.Println(count)
	// fmt.Println(err)
	// fmt.Println(string(buf))

	// var temp interface{} = new(interface{})

	news := NewsInfo{}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&news)
	if err != nil {
		fmt.Println("decode", err.Error())
		return
	}

	//fmt.Printf("%#v", news)

	for _, item := range news.Articles {
		fmt.Println(item.Title)
	}
}
