package main

import (
	"github.com/gin-gonic/gin"
	"go-app/routes"
)

func main() {
	r := gin.Default()
	
	routes.Routes(r)

	r.Run(":8080")
}
