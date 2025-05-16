package repo

import (
	"BlogHub/pkg/db"
	"BlogHub/pkg/models"
	"fmt"
)

func SaveBlog(blog models.Blog) (models.Blog, error) {
	err := db.DB.Save(&blog).Error
	if err != nil {
		return blog, err
	}

	return blog, nil
}

func FindAllBlogs() ([]models.Blog, error) {
	var blogs []models.Blog
	if err := db.DB.Find(&blogs).Error; err != nil {

		return blogs, err
	}
	return blogs, nil
}

func FindBlogByID(blogID string) (models.Blog, error) {
	var blog models.Blog
	if err := db.DB.Where("id=?", blogID).First(&blog).Error; err != nil {
		return blog, err
	}
	if blog.ID == 0 {
		return blog, fmt.Errorf("record not found")
	}
	return blog, nil
}

func DeleteBlogByID(blog models.Blog) error {
	if err := db.DB.Delete(&blog).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBlog(blog models.Blog) (models.Blog, error) {
	if err := db.DB.Save(&blog).Error; err != nil {
		return models.Blog{}, err
	}

	return blog, nil
}
