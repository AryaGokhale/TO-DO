package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AryaGokhale/todo/api/models"
	"github.com/AryaGokhale/todo/api/token"
	"github.com/gin-gonic/gin"
)

var users = []models.User{}

func (server *Server) SignupUser(c *gin.Context) {

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

func (server *Server) LoginUser(c *gin.Context) {

	var loggedUser models.User

	err := json.NewDecoder(c.Request.Body).Decode(&loggedUser)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	for _, u := range users {
		if u.Email == loggedUser.Email && u.Password == loggedUser.Password {

			fmt.Println("Login successfull")

			validToken, err := token.GenerateJWT(u.Name)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
				return
			}

			sessionID := fmt.Sprintf("%d", time.Now().UnixNano())

			c.IndentedJSON(http.StatusOK, gin.H{
				"token":     validToken,
				"sessionID": sessionID,
			})

		}
	}

}
