package main

import (
	"fmt"
	spaceflight "iu-go-homework/spaceflight"
	"net/http"
	"os"
	"sync"
)

type OrderArticle struct {
	Page int
	spaceflight.Article
}

/*
func NewOrderArticle(p int, a spaceflight.Article) OrderArticle {
	return OrderArticle{

	}
}
*/

func main() {
	query := ""
	if len(os.Args) > 1 {
		query = os.Args[1]
	}
	client := spaceflight.NewClient(
		"https://api.spaceflightnewsapi.net/v3",
		os.Getenv("DATA_PROVIDER_APIKEY"),
	)

	ch := make(chan OrderArticle)
	wg := &sync.WaitGroup{}

	//from := time.Now().AddDate(0, -1, 0)
	//fmt.Printf("%#v\n", from.Format("2006-01-02"))

	const articlesCount = 3
	const pageCount = 4

	for i := 0; i < pageCount; i++ {
		wg.Add(1)
		go func(reqi int, channel2 chan OrderArticle) {
			defer wg.Done()
			var news = spaceflight.NewsList{}
			var params = spaceflight.UrlParams{
				"_limit":         fmt.Sprint(articlesCount),
				"_start":         fmt.Sprint(reqi * articlesCount),
				"title_contains": query,
				//"PublishedAt_gt": from.Format("2006-01-02"),
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

			fmt.Printf("\n=== Loading page %d ===\n", reqi)

			for _, item := range news.Articles {
				channel2 <- OrderArticle{reqi, item}
			}
		}(i, ch)
	}
	orders := map[int][]spaceflight.Article{}
	go func() {
		for el := range ch {
			orders[el.Page] = append(orders[el.Page], el.Article)
			//fmt.Println(el.Title)
		}
	}()

	fmt.Println("wait go-s")
	wg.Wait()
	fmt.Println("wait done")
	close(ch)
	for i := 0; i < pageCount; i++ {
		fmt.Printf("\n=== Page %d ===\n", i)
		for _, v := range orders[i] {
			fmt.Println(v.Title)
		}
	}
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
