package router

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/config"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/constant"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/middleware"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/log"
)

var router *gin.Engine

func StartServer() {
	router = gin.Default()
	
	router.Use(cors.New(cors.Config{
		AllowOrigins:     config.GetInstance().Server.Whitelist,
		AllowCredentials: true,
	}))

	router.Use(gin.CustomRecovery(middleware.PanicHandler))

	RegisterRoutes()

	port := config.GetInstance().Server.Port
	err := router.Run(":" + port)

	log.Print(constant.TAG_ROUTER, fmt.Sprintf("Starting Server on port %s", port))
	if err != nil {
		log.Fatal(constant.TAG_ROUTER, err, "Failed to start the server")
	}
}
