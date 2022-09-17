package model

type (
	Book struct {
		ID     int    `json:"id"`
		Title  string `json:"title" form:"title"`
		Author string `json:"author" form:"author"`
	}
)
