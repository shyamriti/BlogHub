package dto

type ExcelResp struct {
	Blogs   []CreateBlogReq `json:"blogs"`
	Message string        `json:"message"`
}
