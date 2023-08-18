package router

import (
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/handler"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/middleware"
)

func RegisterRoutes() {
	AuthRoutes()
	TodoRoutes()
}

func AuthRoutes() {
	authRouter := router.Group(AUTH_BASE_URL)

	authRouter.POST("/registration", handler.Registration)
	authRouter.POST("/login", handler.Login)

	authRouter.Use(middleware.RequireAuth)
	authRouter.PATCH("/change-password", handler.ChangePassword)
	authRouter.POST("/logout", handler.Logout)
}

func TodoRoutes() {
	todoRoutes := router.Group(TODO_BASE_URL)
	todoRoutes.Use(middleware.RequireAuth)

	todoRoutes.GET("", handler.GetAllTodo)
	todoRoutes.POST("", handler.CreateTodo)
	todoRoutes.PATCH("/:todoId", handler.UpdateTodo)
	todoRoutes.PATCH("/delete/:todoId", handler.DeleteTodo)
}
