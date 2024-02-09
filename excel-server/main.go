package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users = []User{
	{ID: "1", Username: "user1", Email: "user1@example.com"},
	{ID: "2", Username: "user2", Email: "user2@example.com"},
}

func excel(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile)

	// Set headers for file download
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=users.xlsx")

	// Serve file
	http.ServeFile(w, r, tempFile)
}

func main() {
	http.HandleFunc("/user", excel)
	http.ListenAndServe(":8080", nil)
}
