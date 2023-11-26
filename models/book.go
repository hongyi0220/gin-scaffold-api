package models

// Book represents a book structure
type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

// Mock database
var Books = []Book{
    {ID: "1", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
    // More books...
}