package model

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"user"`
}
