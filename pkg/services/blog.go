package services

import (
	"BlogHub/pkg/dto"
	"BlogHub/pkg/models"
	"BlogHub/pkg/repo"
	"BlogHub/pkg/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ListOfAllBlogs() ([]models.Blog, error) {
	blogs, err := repo.FindAllBlogs()
	return blogs, err
}

func CreateBlogService(reqBlog dto.CreateBlogReq) (dto.BlogResponse, error) {
	var err error
	blog := models.Blog{
		Title:   reqBlog.Title,
		Caption: reqBlog.Caption,
		UserID:  reqBlog.UserID,
	}
	blog, err = repo.SaveBlog(blog)

	if err != nil {
		return dto.BlogResponse{}, err
	}

	return dto.BlogResponse{
		ID:      blog.ID,
		Title:   blog.Title,
		Caption: blog.Caption,
		// Name:    blog.User.Name,
	}, nil
}

func ListOfAllBlogByID(blogID string) (dto.BlogResponse, error) {
	blog, err := repo.FindBlogByID(blogID)
	if err != nil {
		return dto.BlogResponse{}, err
	}

	userName, err := repo.FindUserNameByID(blog.UserID)
	if err != nil {
		return dto.BlogResponse{}, err
	}

	return dto.BlogResponse{
		Title:     blog.Title,
		Caption:   blog.Caption,
		Name:      userName,
		CreatedAt: blog.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func DeleteBlogByID(c *gin.Context, blogID string) error {
	blog, err := repo.FindBlogByID(blogID)
	if err != nil {
		return err
	}

	if blog.UserID != utils.GetUserID(c) {
		// c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": ""})
		return fmt.Errorf("you are not authorized to delete this Blog")
	}

	if err = repo.DeleteBlogByID(blog); err != nil {
		return err
	}
	return nil
}

func UpdateBlog(c *gin.Context, blogID string, reqBlog dto.UpdateBlogReq) (dto.BlogResponse, error) {
	userID := utils.GetUserID(c)
	blog, err := repo.FindBlogByID(blogID)
	if err != nil {
		return dto.BlogResponse{}, err
	}
	if blog.UserID != userID {
		return dto.BlogResponse{}, fmt.Errorf("you are not authorized to delete this Blog")
	}

	blog.Title = reqBlog.Title
	blog.Caption = reqBlog.Caption

	blog, err = repo.UpdateBlog(blog)
	if err != nil {
		return dto.BlogResponse{}, err
	}
	userName, err := repo.FindUserNameByID(blog.UserID)
	if err != nil {
		return dto.BlogResponse{}, err
	}

	return dto.BlogResponse{
		ID:      blog.ID,
		Title:   blog.Title,
		Caption: blog.Caption,
		Name:    userName,
		CreatedAt: blog.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
