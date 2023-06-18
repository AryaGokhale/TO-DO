package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func (server *Server) Initialize() {

	server.Router = gin.Default()
	server.initializeRoutes()

}

func (server *Server) Run(address string) {

	fmt.Println("Listning on port 8080")
	log.Fatal(http.ListenAndServe(address, server.Router))
}
