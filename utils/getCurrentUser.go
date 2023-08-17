package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/models"
)

func GetCurrentUser(ctx *gin.Context) *models.User {
	user, _ := ctx.Get("user")
	return user.(*models.User)
}
