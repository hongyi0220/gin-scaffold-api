package main

import (
    "github.com/gin-gonic/gin"
    // "net/http"
	"github.com/hongyi0220/gin-scaffold-api/routes"
)

// Book represents a book structure
type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

var books = []Book{
    {ID: "1", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
    // Add more books here
}

func main() {
    router := gin.Default()
	
	// Setup routes
    routes.SetupRoutes(router)
    

    // Start server
    router.Run(":8080")
}

