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
var ID uint32

type NoteReq struct {
	SID  string `json:"sid"`
	NOTE string `json:"note"`
}

func prepareNote(note, author string) models.Note {

	ID = ID + 1
	newNote := models.Note{ID: ID, Content: note, Author: author}
	return newNote
}

func (server *Server) CreateNote(c *gin.Context) {

	var noteReq NoteReq

	err := json.NewDecoder(c.Request.Body).Decode(&noteReq)

	if err != nil {
		return
	}

	token, err := jwt.Parse(noteReq.SID, func(token *jwt.Token) (interface{}, error) {
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

	author := claims["username"].(string)
	note := noteReq.NOTE
	newNote := prepareNote(note, author)

	notes = append(notes, newNote)
	fmt.Println("Note created successfull")
	c.IndentedJSON(http.StatusOK, newNote.ID)

}

type readNote struct {
	ID   uint32 `json:"id"`
	NOTE string `json:"note"`
}

type readNoteResponse struct {
	Notes []readNote `json:"notes"`
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

	//userNotes := []models.Note{}
	readNotes := []readNote{}

	for _, note := range notes {
		if note.Author == name {

			_note := readNote{ID: note.ID, NOTE: note.Content}
			readNotes = append(readNotes, _note)
			fmt.Println("Content is: ", note.Content)
		}
	}

	data := readNoteResponse{
		Notes: readNotes,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	c.IndentedJSON(http.StatusOK, jsonData)

}

type deleteReq struct {
	SID string `json:"sid"`
	ID  uint32 `json:"id"`
}

func (server *Server) DeleteNote(c *gin.Context) {

	var req deleteReq

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
