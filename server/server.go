package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"

	"github.com/AryaGokhale/todo/models"
)

type Server struct {
	Router  *gin.Engine
	Session *sessions.CookieStore
}

var server = Server{}

var users = []models.User{}

var notes = []models.Note{}

func signupUser(c *gin.Context) {

	var newUser models.User

	err := c.BindJSON(&newUser)

	if err != nil {

		c.IndentedJSON(http.StatusBadRequest, newUser)
		return
	} else {

		users = append(users, newUser)
		c.IndentedJSON(http.StatusOK, newUser)
		fmt.Println("User created successfull")
	}

}

func loginUser(c *gin.Context) {

	var loggedUser models.User

	var sessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

	err := json.NewDecoder(c.Request.Body).Decode(&loggedUser)

	session, _ := sessionStore.Get(c.Request, "session-name")

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	for _, u := range users {
		if u.Email == loggedUser.Email && u.Password == loggedUser.Password {

			fmt.Println("Login successfull")

			session.Values["isLoggedIn"] = true

			c.IndentedJSON(http.StatusFound, session.ID)

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
	fmt.Println("Note created successfull")

}

func readNote(c *gin.Context) {

	userNotes := []models.Note{}

	var username models.Note

	err := json.NewDecoder(c.Request.Body).Decode(&username)

	if err != nil {

		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return

	}
	for _, note := range notes {
		if note.Author == username.Author {

			userNotes = append(userNotes, note)
			fmt.Println("Content is: ", note.Content)
		}
	}

	c.IndentedJSON(http.StatusOK, userNotes)

}

func deleteNote(c *gin.Context) {

	var req models.Note

	err := json.NewDecoder(c.Request.Body).Decode(&req)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	noteIndex := -1

	for i, note := range notes {

		if note.ID == req.ID {

			noteIndex = i
		}
	}

	notes = append(notes[:noteIndex], notes[noteIndex+1:]...)

	fmt.Println("Deleted note successfully")

}

func (s *Server) Runserver() {
	s.Router = gin.Default()

	s.Router.POST("/signup", signupUser)
	s.Router.POST("/login", loginUser)
	s.Router.POST("/notes", createNote)
	s.Router.GET("/notes", readNote)
	s.Router.DELETE("/notes", deleteNote)

	s.Router.Run(os.Getenv("SERVER_PORT"))
}
