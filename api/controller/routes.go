package controller

func (server *Server) initializeRoutes() {

	server.Router.POST("/signup", server.SignupUser)
	server.Router.POST("/login", server.LoginUser)
	server.Router.POST("/notes", server.CreateNote)
	server.Router.GET("/notes", server.ReadNote)
	server.Router.DELETE("/notes", server.DeleteNote)
}
