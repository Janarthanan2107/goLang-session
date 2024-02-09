package controllers

import (
	"net/http"
	"os"
	"strconv"
	"user/models"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
)

var users = []models.User{

	{
		ID:       "1",
		Username: "Bret",
		Email:    "Sincere@april.biz",
	},
	{
		ID:       "2",
		Username: "Antonette",
		Email:    "Shanna@melissa.tv",
	},
	{
		ID:       "3",
		Username: "Samantha",
		Email:    "Nathan@yesenia.net",
	},
	{
		ID:       "4",
		Username: "Karianne",
		Email:    "Julianne.OConner@kory.org",
	},
	{
		ID:       "5",
		Username: "Kamren",
		Email:    "Lucio_Hettinger@annie.ca",
	},
	{
		ID:       "6",
		Username: "Leopoldo_Corkery",
		Email:    "Karley_Dach@jasper.info",
	},
	{
		ID:       "7",
		Username: "Elwyn.Skiles",
		Email:    "Telly.Hoeger@billy.biz",
	},
	{
		ID:       "8",
		Username: "Maxime_Nienow",
		Email:    "Sherwood@rosamond.me",
	},
	{
		ID:       "9",
		Username: "Delphine",
		Email:    "Chaim_McDermott@dana.io",
	},
	{
		ID:       "10",
		Username: "Moriah.Stanton",
		Email:    "Rey.Padberg@karina.biz",
	},
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func GetUsersExcel(c *gin.Context) {
	file := excelize.NewFile()

	// Create a new sheet
	index := file.NewSheet("Users")

	// Set headers
	headers := []string{"ID", "Username", "Email"}
	for col, header := range headers {
		cell := excelize.ToAlphaString(col) + "1"
		file.SetCellValue("Users", cell, header)
	}

	// Set data
	for row, user := range users {
		file.SetCellValue("Users", "A"+strconv.Itoa(row+2), user.ID)
		file.SetCellValue("Users", "B"+strconv.Itoa(row+2), user.Username)
		file.SetCellValue("Users", "C"+strconv.Itoa(row+2), user.Email)
	}

	// Set active sheet of the workbook
	file.SetActiveSheet(index)

	// Save the XLSX file to a temporary location
	tempFile := "users.xlsx"
	if err := file.SaveAs(tempFile); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile)

	// Set headers for file download
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=users.xlsx")

	// Stream the file to the client
	c.File(tempFile)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Generate a unique ID for the new user (You might use UUID or any other method)
	newUser.ID = "3" // For simplicity, manually setting ID to "3"
	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, user := range users {
		if user.ID == id {
			// Update user details
			users[i].Username = updatedUser.Username
			users[i].Email = updatedUser.Email
			c.JSON(http.StatusOK, users[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	for i, user := range users {
		if user.ID == id {
			// Remove user from the slice
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
