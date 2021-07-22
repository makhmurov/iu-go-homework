package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("hi")
	client := NewClient("https://newsapi.org/v2", os.Getenv("DATA_PROVIDER_APIKEY"))

	var params = UrlParams{
		"q":        "tesla",
		"form":     "2021-06-21",
		"sortBy":   "publishedAt",
		"pageSize": "5",
		"language": "ru",
	}

	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/everything?%s", client.BaseUrl, params.Query()),
		nil,
	)

	_ = err

	news := NewsInfo{}

	err = client.SendRequest(request, &news)
	if err != nil {
		fmt.Printf("Error %s", err.Error())
		return
	}

	for _, item := range news.Articles {
		fmt.Println(item.Title)
	}
}
