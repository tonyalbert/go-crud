package Handlers

import (
	"go-app/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeHandler View Handlers
func HomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// IndexHandler API Handlers
func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func GetProducts(c *gin.Context) {
	database.InitDB()

	rows, err := database.DB.Query("SELECT id, name, price FROM products")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro":    err.Error(),
			"message": "Error getting products",
		})
		return
	}

	defer rows.Close()

	var products []struct {
		Id    int
		Name  string
		Price float64
	}

	for rows.Next() {
		var product struct {
			Id    int
			Name  string
			Price float64
		}

		if err := rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "Error getting products",
			})
			return
		}

		products = append(products, product)
	}

	c.JSON(http.StatusOK, gin.H{"products": products})

	defer database.CloseDB()
}

func SaveProduct(c *gin.Context) {
	var requestBody struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}

	// Faça o bind do corpo JSON para a estrutura
	if err := c.BindJSON(&requestBody); err != nil {
		// Se houver um erro no bind, envie uma resposta de erro
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
		return
	}

	database.InitDB()

	_, err := database.DB.Exec("INSERT INTO products (name, price) VALUES (?, ?)", requestBody.Name, requestBody.Price)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Error saving product",
		})
		return
	}

	defer database.CloseDB()

	c.JSON(http.StatusOK, gin.H{"message": "Product saved successfully", "product": requestBody})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	database.InitDB()

	_, err := database.DB.Exec("DELETE FROM products WHERE id = ?", id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Error deleting product",
		})
		return
	}

	defer database.CloseDB()

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
