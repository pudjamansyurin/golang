package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pudjamansyurin/gin-gorm/controllers"
	"github.com/pudjamansyurin/gin-gorm/models"
)

func main() {
	r := gin.Default()

	models.ConnectDB()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PUT("/books/:id", controllers.UpdateBook)

	r.Run()
}
