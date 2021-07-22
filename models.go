package main

import "time"

type Source struct {
	ID   interface{} `json:"id"`
	Name string      `json:"name"`
}

type Article struct {
	Source      Source      `json:"source"`
	Author      interface{} `json:"author"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	URL         string      `json:"url"`
	URLToImage  string      `json:"urlToImage"`
	PublishedAt time.Time   `json:"publishedAt"`
	Content     string      `json:"content"`
}

type NewsInfo struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}
