package main

import (
	authcontrollers "gin/controllers/authControllers"
	productcontrollers "gin/controllers/productControllers"
	"gin/database"
	"gin/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	database.ConnectToDB()

	// Create a Gin router
	r := gin.Default()

	// Auth routes
	r.POST("/login", authcontrollers.Login)
	r.POST("/register", authcontrollers.Register)
	r.GET("/logout", authcontrollers.Logout)

	// Protected routes
	api := r.Group("/api")
	api.Use(middleware.JWTMiddleware()) // JWT middleware diaktifkan untuk semua route dalam /api
	{
		api.GET("/products", productcontrollers.GetAll)
		api.GET("/products/:id", productcontrollers.GetByID)
		api.POST("/products", productcontrollers.Create)
		api.PUT("/products/:id", productcontrollers.Update)
		api.DELETE("/products/:id", productcontrollers.Delete)
	}

	// Start the server
	r.Run(":8080")
}
