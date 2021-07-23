package spaceflight

import "time"

type ApiErrorDoc struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ApiError struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}

type NewsSite struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Articles  []string `json:"articles"`
	Blogs     []string `json:"blogs"`
	Reports   []string `json:"reports"`
	CreatedBy string   `json:"created_by"`
	UpdatedBy string   `json:"updated_by"`
}

type Launch struct {
	ID        string   `json:"id"`
	LaunchID  string   `json:"launchId"`
	Name      string   `json:"name"`
	Articles  []string `json:"articles"`
	Blogs     []string `json:"blogs"`
	Provider  string   `json:"provider"`
	CreatedBy string   `json:"created_by"`
	UpdatedBy string   `json:"updated_by"`
}

type Event struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	EventID   int      `json:"eventId"`
	Articles  []string `json:"articles"`
	Blogs     []string `json:"blogs"`
	Provider  string   `json:"provider"`
	CreatedBy string   `json:"created_by"`
	UpdatedBy string   `json:"updated_by"`
}

type Article struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	ImageURL    string    `json:"imageUrl"`
	Summary     string    `json:"summary"`
	PublishedAt time.Time `json:"publishedAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Featured    bool      `json:"featured"`
	NewsSite    string    `json:"newsSite"`
	Launches    []Launch  `json:"launches"`
	Events      []Event   `json:"events"`
}

type NewsList struct {
	Articles []Article `json:"articles"`
}
