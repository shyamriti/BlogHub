package controllers

import (
	"BlogHub/pkg/dto"
	"BlogHub/pkg/services"
	"BlogHub/pkg/utils"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBlog(c *gin.Context) {
	var reqBlog dto.CreateBlogReq

	if err := c.ShouldBindJSON(&reqBlog); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	reqBlog.UserID = utils.GetUserID(c)

	respBlog, err := services.CreateBlogService(reqBlog)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"db_error": "Failed to insert Blog",
			"details":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Blog created successfully",
		"Blog":    respBlog,
	})
}

func GetBlogs(c *gin.Context) {

	blogs, err := services.ListOfAllBlogs()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Blogs": blogs,
	})
}

func GetBlogByBlogId(c *gin.Context) {

	blogID := c.Param("id")

	blog, err := services.ListOfAllBlogByID(blogID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, blog)
}

func DeleteBlog(c *gin.Context) {
	blogID := c.Param("id")

	if err := services.DeleteBlogByID(c, blogID); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error":   "Failed to delete Blog",
			"details": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

func UpdateBlog(c *gin.Context) {
	var reqBlog dto.UpdateBlogReq
	id := c.Param("id")

	if err := c.ShouldBindJSON(&reqBlog); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error":   "Failed to Bind Data",
			"details": err.Error(),
		})
		return
	}

	blog, err := services.UpdateBlog(c, id, reqBlog)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "failed during update blog",
			"details": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "Blog Updated successfully",
		"Blog":    &blog,
	})
}

func CreateMultipleBlog(c *gin.Context) {
	var reqBlogs []dto.CreateBlogReq
	var respBlogs []dto.BlogResponse

	resp, err := http.Get("http://localhost:8080/excel")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "failed to call /excel",
			"details": err.Error(),
		})
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var ExcelResp dto.ExcelResp

	if err = json.Unmarshal(body, &ExcelResp); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse /excel response", "details": err.Error()})
	}

	reqBlogs = ExcelResp.Blogs

	for i:=0; i< len(reqBlogs); i++{
		respBlog, err := services.CreateBlogService(reqBlogs[i])
		if err!=nil{
			c.AbortWithStatusJSON(400, gin.H{
			"db_error": "Failed to insert Blog",
			"details":  err.Error(),
		})
		return
		}
		respBlogs = append(respBlogs, respBlog)
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "/add executed and /excel was also called",
		"blogs": respBlogs,
	})
}
