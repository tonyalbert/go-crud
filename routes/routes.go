package routes

import (
	"go-app/handlers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	// Load HTML templates
	r.LoadHTMLFiles("templates/index.html")

	// VIEW
	r.GET("/", Handlers.HomeHandler)

	// API
	r.GET("/api", Handlers.IndexHandler)
	r.GET("/api/products", Handlers.GetProducts)
	r.POST("/api/product", Handlers.SaveProduct)
	r.DELETE("/api/product/:id", Handlers.DeleteProduct)
}