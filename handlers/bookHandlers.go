package handlers

import (
    "net/http"
	"github.com/hongyi0220/gin-scaffold-api/models"
    "github.com/gin-gonic/gin"
)

// GetBooks responds with the list of all books as JSON.
func GetBooks(c *gin.Context) {
    c.JSON(http.StatusOK, models.Books)
}

// GetBookByID locates the book whose ID value matches the id
// parameter sent by the client, then returns that book as a response.
func GetBookByID(c *gin.Context) {
    id := c.Param("id")

    for _, a := range models.Books {
        if a.ID == id {
            c.JSON(http.StatusOK, a)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// PostBook adds a new book from JSON received in the request body.
func PostBook(c *gin.Context) {
    var newBook models.Book

    // Call BindJSON to bind the received JSON to newBook.
    if err := c.BindJSON(&newBook); err != nil {
        return
    }

    // Add the new book to the slice.
    models.Books = append(models.Books, newBook)
    c.JSON(http.StatusCreated, newBook)
}

// UpdateBook updates the book whose ID value matches the id
// parameter sent by the client.
func UpdateBook(c *gin.Context) {
    id := c.Param("id")
    var updatedBook models.Book

    if err := c.BindJSON(&updatedBook); err != nil {
        return
    }

    for i, a := range models.Books {
        if a.ID == id {
            models.Books[i] = updatedBook
            c.JSON(http.StatusOK, updatedBook)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// DeleteBook removes the book whose ID value matches the id
// parameter sent by the client.
func DeleteBook(c *gin.Context) {
    id := c.Param("id")

    for i, a := range models.Books {
        if a.ID == id {
            models.Books = append(models.Books[:i], models.Books[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}