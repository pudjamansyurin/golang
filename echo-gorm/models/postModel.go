package models

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Likes   int    `json:"likes"`
	Author  User   `json:"author"`
}
