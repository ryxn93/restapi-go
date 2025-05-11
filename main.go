package main

import (
	productcontrollers "gin/controllers/productControllers"
	"gin/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectToDB()
	router := gin.Default()
	router.POST("/products", productcontrollers.Create)
	router.Run()
}
