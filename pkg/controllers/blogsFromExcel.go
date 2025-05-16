package controllers

import (
	"BlogHub/pkg/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

func GetDataFromExcel(c *gin.Context) {
	var reqBlogs []dto.CreateBlogReq

	fileName := "blog.xlsx"

	wb, err := xlsx.OpenFile(fileName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "file not open",
			"details": err.Error(),
		})
		return
	}

	sh, ok := wb.Sheet["Blogs"]
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Sheet not found",
		})
		return
	}

	for i := 1; i < len(sh.Rows); i++ {
		row := sh.Rows[i]
		if len(row.Cells) < 3 {
			continue 
		}

		title := row.Cells[0].String()
		caption := row.Cells[1].String()
		userID, err := row.Cells[2].Int()
		if err != nil {
			continue
		}

		reqBlog := dto.CreateBlogReq{
			Title:   title,
			Caption: caption,
			UserID:  uint(userID),
		}
		reqBlogs = append(reqBlogs, reqBlog)
	}

	var resp dto.ExcelResp

	resp.Blogs = reqBlogs

	c.JSON(http.StatusOK, resp)
}
