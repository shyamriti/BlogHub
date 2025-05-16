package routers

import (
	"BlogHub/pkg/controllers"
	"BlogHub/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()

	
	r.GET("/hello", controllers.Hello)
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/add", controllers.CreateBlog)
		authorized.GET("/blogs", controllers.GetBlogs)
		authorized.GET("/blog/:id", controllers.GetBlogByBlogId)
		authorized.DELETE("/delete/:id", controllers.DeleteBlog)
		authorized.PUT("/update/:id", controllers.UpdateBlog)
	}

	return r
}
