package hackernews

import "time"

type Story struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

// #######

type Source struct {
	ID   interface{} `json:"id"`
	Name string      `json:"name"`
}

type Article struct {
	Source Source `json:"source"`
	Author      interface{} `json:"author"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	URL         string      `json:"url"`
	URLToImage  string      `json:"urlToImage"`
	PublishedAt time.Time   `json:"publishedAt"`
	Content     string      `json:"content"`
}

type NewsInfo struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []Article `json:"articles"`
}