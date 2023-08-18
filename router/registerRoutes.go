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
	authRouter.PATCH("change-password", middleware.RequireAuth, handler.ChangePassword)
	authRouter.POST("/logout", middleware.RequireAuth, handler.Logout)
}

func TodoRoutes() {
	todoRoutes := router.Group(TODO_BASE_URL)
	todoRoutes.GET("", middleware.RequireAuth, handler.GetAllTodo)
	todoRoutes.POST("", middleware.RequireAuth, handler.CreateTodo)
	todoRoutes.PATCH("/:todoId", middleware.RequireAuth, handler.UpdateTodo)
	todoRoutes.PATCH("/delete/:todoId", middleware.RequireAuth, handler.DeleteTodo)
}
