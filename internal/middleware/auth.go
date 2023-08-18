package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/config"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/constant"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/repository"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/log"
)

func RequireAuth(ctx *gin.Context) {
	signedToken, err := ctx.Cookie(constant.ACCESS_TOKEN)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	rawToken, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.GetInstance().JWT.SecretKey), nil
	})

	if err != nil {
		log.Error(constant.TAG_MIDDLEWARE, rawToken, err, "auth[RequireAuth]: jwt.Parse failed to parse signedToken")
	}

	if claims, ok := rawToken.Claims.(jwt.MapClaims); ok && rawToken.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		uid := claims["uid"].(string)

		user := repository.FindUserByUid(uid)
		if user.Email == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		ctx.Set("user", user)
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
