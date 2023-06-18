package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"github.com/AryaGokhale/todo/models"
)

type Token struct {
	Username    string `json:"name"`
	TokenString string `json:"token"`
}

var users = []models.User{}

var notes = []models.Note{}

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

func generateJWT(username string) (string, error) {

	var mySigningKey = []byte(jwtKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username

	//claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong which signing token: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

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

// func loginUser(c *gin.Context) {

// 	var loggedUser models.User

// 	err := json.NewDecoder(c.Request.Body).Decode(&loggedUser)

// 	if err != nil {
// 		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	for _, u := range users {
// 		if u.Email == loggedUser.Email && u.Password == loggedUser.Password {

// 			fmt.Println("Login successfull")
// 			c.IndentedJSON(http.StatusFound, loggedUser)

// 		}
// 	}

// }

func loginUser(c *gin.Context) {

	var loggedUser models.User

	err := json.NewDecoder(c.Request.Body).Decode(&loggedUser)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	for _, u := range users {
		if u.Email == loggedUser.Email && u.Password == loggedUser.Password {

			fmt.Println("Login successfull")

			//sessionID := uuid.New().String()

			validToken, err := generateJWT(u.Name)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
				return
			}

			sessionID := fmt.Sprintf("%d", time.Now().UnixNano())

			c.IndentedJSON(http.StatusOK, gin.H{
				"token":     validToken,
				"sessionID": sessionID,
			})

			//c.IndentedJSON(http.StatusFound, sessionID)

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

	//var req models.Note

	//err := json.NewDecoder(c.Request.Body).Decode(&req)

	// if err != nil {

	// 	http.Error(c.Writer, err.Error(), http.StatusBadRequest)
	// 	return

	// }

	for _, note := range notes {
		if note.Author == name {

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

func main() {
	router := gin.Default()

	router.POST("/signup", signupUser)
	router.POST("/login", loginUser)
	router.POST("/notes", createNote)
	router.GET("/notes", readNote)
	router.DELETE("/notes", deleteNote)
	router.Run("localhost:8080")
}
