package main

import (
	"fmt"
	spaceflight "iu-go-homework/spaceflight"
	"net/http"
	"os"
	"sync"
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

	ch := make(chan spaceflight.Article)
	wg := &sync.WaitGroup{}

	const articlesCount = 3
	const pageCount = 4

	for i := 0; i < pageCount; i++ {
		wg.Add(1)
		go func(reqi int, channel2 chan spaceflight.Article) {
			defer wg.Done()
			var news = spaceflight.NewsList{}
			var params = spaceflight.UrlParams{
				"_limit":         fmt.Sprint(articlesCount),
				"_start":         fmt.Sprint(reqi * articlesCount),
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

			err = client.SendRequest(request, &news.Articles)

			if err != nil {
				fmt.Printf("Error %s", err.Error())
				return
			}

			fmt.Printf("\n=== Page %d ===\n", reqi)

			for _, item := range news.Articles {
				channel2 <- item
			}
		}(i, ch)
	}

	go func() {
		for article := range ch {
			fmt.Println(article.Title)
		}
	}()

	fmt.Println("wait go-s")
	wg.Wait()
	fmt.Println("wait done")
	close(ch)
	fmt.Println("Program stop")

	//

	/*
		if err != nil {
			fmt.Printf("Error %s", err.Error())
			return
		}

		for _, item := range news.Articles {
			fmt.Println(item.Title)
		}
	*/

	// Get article counts
	// Get article short list

	// Get command
	// Get article
}
