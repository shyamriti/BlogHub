package dto

type CreateBlogReq struct {
	Title   string `json:"title" binding:"required,min=3,max=100"`
	Caption string `json:"caption" binding:"required,min=3,max=10"`
	UserID  uint
}

type UpdateBlogReq struct {
	ID        uint   `json:"id"`
	Title     string `json:"title" binding:"required,min=3,max=100"`
	Caption   string `json:"caption" binding:"required,min=3,max=10"`
	CreatedAt string `json:"created_at"`
}
