package model

type CreateBlogRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateBlogRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
