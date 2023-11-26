package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/hongyi0220/gin-scaffold-api/handlers"
)

// SetupRoutes configures all the routes for the application.
func SetupRoutes(router *gin.Engine) {
    router.GET("/books", handlers.GetBooks)
    router.GET("/books/:id", handlers.GetBookByID)
    router.POST("/books", handlers.PostBook)
    router.PUT("/books/:id", handlers.UpdateBook)
    router.DELETE("/books/:id", handlers.DeleteBook)

    // Add more routes here
}