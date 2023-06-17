package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AryaGokhale/todo/models"
)

var users = []models.User{}

var notes = []models.Note{}

func signupUser(c *gin.Context) {

	var newUser models.User

	err := c.BindJSON(&newUser)

	if err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
	fmt.Println("User created successfull")
}

func loginUser(c *gin.Context) {

	var loggedUser models.User

	err := json.NewDecoder(c.Request.Body).Decode(&loggedUser)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	for _, u := range users {
		if u.Email == loggedUser.Email && u.Password == u.Password {
			c.IndentedJSON(http.StatusFound, loggedUser)
			fmt.Println("Login successfull")
		}
	}

}

func createNote(c *gin.Context) {

	var newNote models.Note

	err := json.NewDecoder(c.Request.Body).Decode(&newNote)

	if err != nil {
		return
	}

	notes = append(notes, newNote)
	//c.IndentedJSON(http.StatusCreated, newNote)
	fmt.Println("Note created successfull")

}

func readNote(c *gin.Context) {

	var username string

	userNotes := []models.Note{}

	username = c.Request.URL.Query().Get("author")

	for _, note := range notes {
		if note.Author == username {

			userNotes = append(userNotes, note)
			fmt.Println("Content is: ", note.Content)
		}
	}

}

func main() {
	router := gin.Default()

	router.POST("/signup", signupUser)
	router.POST("/login", loginUser)
	router.POST("/notes", createNote)
	router.GET("/notes", readNote)
	router.Run("localhost:8080")
}
