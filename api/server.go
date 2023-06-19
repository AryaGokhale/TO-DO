package api

import "github.com/AryaGokhale/todo/api/controller"

var server = controller.Server{}

func Run() {

	server.Initialize()
	server.Run(":8080")
}
