package main

import (
	"fmt"
	spaceflight "iu-go-homework/spaceflight"
	"net/http"
	"os"
)

func main() {
	query := ""
	if len(os.Args) > 1 {
		query = os.Args[1]
	}
	client := spaceflight.NewClient(
		"https://api.spaceflightnewsapi.net/v3",
		os.Getenv("DATA_PROVIDER_APIKEY"),
	)

	var params = spaceflight.UrlParams{
		"_limit":         "5",
		"_start":         "0",
		"title_contains": query,
	}

	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/articles?%s", client.BaseUrl, params.Query()),
		nil,
	)

	if err != nil {
		fmt.Printf("Error %s", err.Error())
		return
	}

	var news = spaceflight.NewsList{}

	/*
		sync := make(chan error)
		go func() {

			sync <- err
			close(sync)
		}()
		err = <-sync
	*/
	err = client.SendRequest(request, &news.Articles)
	if err != nil {
		fmt.Printf("Error %s", err.Error())
		return
	}

	for _, item := range news.Articles {
		fmt.Println(item.Title)
	}

	// Get article counts
	// Get article short list

	// Get command
	// Get article
}
