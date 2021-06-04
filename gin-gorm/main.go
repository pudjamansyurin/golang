package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pudjamansyurin/gin-gorm/controllers"
	"github.com/pudjamansyurin/gin-gorm/models"
)

func main() {
	models.ConnectDB()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	v1 := r.Group("/book")
	{
		v1.GET("", controllers.FindBooks)
		v1.GET("/:id", controllers.FindBook)
		v1.POST("", controllers.CreateBook)
		v1.PUT("/:id", controllers.UpdateBook)
		v1.DELETE("/:id", controllers.DeleteBook)
	}

	r.Run()
}
