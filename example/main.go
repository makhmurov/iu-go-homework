package main

import (
	"fmt"
	finder "iu-go-homework"
	"net/http"
	"os"
)

func main() {
	query := "tesla"
	if len(os.Args) > 1 {
		query = os.Args[1]
	}

	fmt.Println("hi")
	client := finder.NewClient("https://newsapi.org/v2", os.Getenv("DATA_PROVIDER_APIKEY"))

	var params = finder.UrlParams{
		"q":        query,
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

	news := finder.NewsInfo{}

	err = client.SendRequest(request, &news)
	if err != nil {
		fmt.Printf("Error %s", err.Error())
		return
	}

	for _, item := range news.Articles {
		fmt.Println(item.Title)
	}
}
