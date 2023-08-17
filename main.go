package main

import (
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/config"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/database"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/router"
)

func main() {
	config.LoadEnvVariables()
	database.Init()
	defer database.CloseConnection()

	router.StartServer()
}
