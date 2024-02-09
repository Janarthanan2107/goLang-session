package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// demo data
var books = []book{
	{ID: "1", Title: "In search of lost time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "It's ends with us", Author: "Marcel Proust", Quantity: 4},
	{ID: "3", Title: "It's starts with us", Author: "Marcel Proust", Quantity: 2},
}

// root function for the server
func root(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Welcome to the server")
}

// getBooks handles GET requests for fetching books
// gin.context ===> used for made accepting the struct for the function
func getBooks(c *gin.Context) {
	// IndentedJSON ===> used for pretty the json we send
	// https.statusOk,StatusNotFound,StatusBadRequest ===> status code 200
	// second parameter for sending data
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	// gin.H ===> is used to write a custom message
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func checkoutBook(c *gin.Context) {
	// getQuery ===> is used to get a guery rised in url
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available."})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
	var newBook book

	// bindJson is a method is what will handle sending the error response
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	found := false

	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			found = true
			break
		}
	}

	if found {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	}
}

func main() {
	router := gin.Default()
	router.GET("/", root)
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)

	router.POST("/books", createBook)
	// NOTE : using curl testing run this command for post
	// @body.json is containing the test json
	// curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"

	router.PATCH("/checkout", checkoutBook)
	// curl localhost:8080/checkout?id=2 --request "PATCH"

	router.DELETE("/books/:id", deleteBook)
	// curl -X DELETE http://localhost:8080/books/1

	// Print a message indicating that the server is running
	fmt.Println("Server is running on the port: http://localhost:8080")

	// Start the server
	router.Run("localhost:8080")
}
