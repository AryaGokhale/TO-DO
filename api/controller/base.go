package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server with a gin Engine field
type Server struct {
	Router *gin.Engine
}

// Initializing the server and the associated routes that it serves
func (server *Server) Initialize() {

	server.Router = gin.Default()
	server.initializeRoutes()

}

// Listen for request on port 8080
func (server *Server) Run(address string) {

	fmt.Println("Listning on port 8080")
	log.Fatal(http.ListenAndServe(address, server.Router))
}
