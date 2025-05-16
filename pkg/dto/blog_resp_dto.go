package dto

type BlogResponse struct {
	ID        uint              `json:"id"`
	Title     string            `json:"title"`
	Caption   string            `json:"caption"`
	Name      string            `json:"name"`
	Comments  []CommentResponse `json:"comments,omitempty"`
	CreatedAt string            `json:"created_at"`
}

type CommentResponse struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	// Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}
