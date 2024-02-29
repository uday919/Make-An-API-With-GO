package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Data Structures and Algorithms", Author: "Durga Kiran", Quantity: 3},
	{ID: "2", Title: "Love Guru", Author: "Venkatesh J", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "SriCharan", Quantity: 8},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}
func checkOutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Missing id query parameter"})
		return
	}
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "No more copies of this book available"})
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}
func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Missing id query parameter"})
		return
	}
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}
func getBookById(id string) (*book, error) {
	for i, book := range books {
		if book.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("No Book Found!")
}
func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkOutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}
