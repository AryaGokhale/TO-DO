package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/AryaGokhale/todo/api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var notes = []models.Note{}
var jwtKey = []byte(os.Getenv("SECRET_KEY"))

func (server *Server) CreateNote(c *gin.Context) {

	var newNote models.Note

	err := json.NewDecoder(c.Request.Body).Decode(&newNote)

	if err != nil {
		return
	}

	notes = append(notes, newNote)
	fmt.Println("Note created successfull")

}

func (server *Server) ReadNote(c *gin.Context) {

	sessionID := c.GetHeader("Authorization")

	if sessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing session ID"})
		return
	}

	token, err := jwt.Parse(sessionID, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session ID"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session ID"})
		return
	}

	name := claims["username"].(string)

	userNotes := []models.Note{}

	for _, note := range notes {
		if note.Author == name {

			userNotes = append(userNotes, note)
			fmt.Println("Content is: ", note.Content)
		}
	}

	c.IndentedJSON(http.StatusOK, userNotes)

}

func (server *Server) DeleteNote(c *gin.Context) {

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
