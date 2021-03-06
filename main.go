package main

import (
	"os"

	"github.com/ervinismu/kejarmimpi/controllers"
	"github.com/ervinismu/kejarmimpi/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	// This function is for AutoMigrate
	// db := configs.InitDb()
	// defer db.Close()

	// var post models.Post
	// var user models.User
	// db.AutoMigrate(&user)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	//This func is for server API
	r := gin.Default()
	//Method for get post and login
	r.GET("/post", controllers.GetPost)
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.Use(middleware.CheckToken)
	//Method for post
	r.POST("/post", controllers.CreatePost)
	r.PUT("/post/:id", controllers.UpdatePost)
	r.DELETE("/post/:id", controllers.DeletePost)
	//Method for logout
	r.GET("/logout", controllers.Logout)
	//Method for profile
	r.GET("/profile/:id", controllers.GetProfile)
	r.PUT("profile/:id", controllers.UpdateProfile)
	r.Run(":" + port)
}
