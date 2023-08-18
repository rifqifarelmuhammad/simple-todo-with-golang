package handler

import (
	"net/http"
	"net/mail"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/config"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/constant"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/dto"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/repository"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/log"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/utils"
	"golang.org/x/crypto/bcrypt"
)

func Registration(ctx *gin.Context) {
	body := dto.RegistrationRequest{}
	err := ctx.Bind(&body)
	if err != nil {
		utils.ResponseHandler(ctx, utils.HTTPResponse{
			ResponseCode:    http.StatusBadRequest,
			ResponseMessage: "Failed to read request body",
			ResponseStatus:  utils.RESPONSE_STATUS_FAILED,
		})

		return
	}

	valid := BodyRequestValidation(ctx, "Registration", body.AuthRequest)
	if !valid {
		return
	}

	if body.ConfirmationPassword == "" {
		utils.ResponseHandler(ctx, utils.HTTPResponse{
			ResponseCode:    http.StatusBadRequest,
			ResponseMessage: "Confirmation password cannot be empty",
			ResponseStatus:  utils.RESPONSE_STATUS_FAILED,
		})

		return
	}

	if body.Password != body.ConfirmationPassword {
		utils.ResponseHandler(ctx, utils.HTTPResponse{
			ResponseCode:    http.StatusBadRequest,
			ResponseMessage: "Password does not match confirmation password",
			ResponseStatus:  utils.RESPONSE_STATUS_FAILED,
		})

		return
	}

	if len(body.Password) < 8 {
		utils.ResponseHandler(ctx, utils.HTTPResponse{
			ResponseCode:    http.StatusBadRequest,
			ResponseMessage: "Password too short",
			ResponseStatus:  utils.RESPONSE_STATUS_FAILED,
		})

		return
	}

	user := repository.FindUserByEmail(body.Email)
	if user.UID != "" {
		utils.ResponseHandler(ctx, utils.HTTPResponse{
			ResponseCode:    http.StatusConflict,
			ResponseMessage: "User already exists",
			ResponseStatus:  utils.RESPONSE_STATUS_FAILED,
		})

		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		log.Error(constant.TAG_SERVICES, hashedPassword, err, "auth[Registration]: bcrypt.GenerateFromPassword failed to hash password")
		panic(err)
	}

	user = repository.CreateUser(body.Email, hashedPassword)

	signedToken := SetAccessToken(ctx, user.UID)

	utils.ResponseHandler(ctx, utils.HTTPResponse{
		ResponseCode:    http.StatusCreated,
		ResponseMessage: "User successfully registered",
		ResponseStatus:  utils.RESPONSE_STATUS_SUCCESS,
		Data:            signedToken,
	})
}

func Login(ctx *gin.Context) {
	body := dto.AuthRequest{}
	err := ctx.Bind(&body)
	if err != nil {
		utils.ResponseHandler(ctx, utils.HTTPResponse{
			ResponseCode:    http.StatusBadRequest,
			ResponseMessage: "Failed to read request body",
			ResponseStatus:  utils.RESPONSE_STATUS_FAILED,
		})

		return
	}

	valid := BodyRequestValidation(ctx, "Registration", body)
	if !valid {
		return
	}

	user := repository.FindUserByEmail(body.Email)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if user.UID == "" || err != nil {
		utils.ResponseHandler(ctx, utils.HTTPResponse{
			ResponseCode:    http.StatusUnauthorized,
			ResponseMessage: "Invalid Email or Password",
			ResponseStatus:  utils.RESPONSE_STATUS_FAILED,
		})

		return
	}

	signedToken := SetAccessToken(ctx, user.UID)

	utils.ResponseHandler(ctx, utils.HTTPResponse{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "Login successful",
		ResponseStatus:  utils.RESPONSE_STATUS_SUCCESS,
		Data:            signedToken,
	})
}

func Logout(ctx *gin.Context) {
	ctx.SetCookie(constant.ACCESS_TOKEN, "", -1, "", "", true, true)
	utils.ResponseHandler(ctx, utils.HTTPResponse{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "Logout successful",
		ResponseStatus:  utils.RESPONSE_STATUS_SUCCESS,
	})
}

func BodyRequestValidation(ctx *gin.Context, funcName string, body dto.AuthRequest) bool {
	if body.Email == "" {
		utils.ResponseHandler(ctx, utils.HTTPResponse{
			ResponseCode:    http.StatusBadRequest,
			ResponseMessage: "Email cannot be empty",
			ResponseStatus:  utils.RESPONSE_STATUS_FAILED,
		})

		return false
	}

	_, err := mail.ParseAddress(body.Email)
	if err != nil {
		utils.ResponseHandler(ctx, utils.HTTPResponse{
			ResponseCode:    http.StatusBadRequest,
			ResponseMessage: "Invalid email",
			ResponseStatus:  utils.RESPONSE_STATUS_FAILED,
		})

		return false
	}

	if body.Password == "" {
		utils.ResponseHandler(ctx, utils.HTTPResponse{
			ResponseCode:    http.StatusBadRequest,
			ResponseMessage: "Password cannot be empty",
			ResponseStatus:  utils.RESPONSE_STATUS_FAILED,
		})

		return false
	}

	return true
}

func SetAccessToken(ctx *gin.Context, uid string) string {
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(time.Hour * 24 * time.Duration(config.GetInstance().JWT.ExpireTime)).Unix(),
	})

	signedToken, err := rawToken.SignedString([]byte(config.GetInstance().JWT.SecretKey))
	if err != nil {
		log.Error(constant.TAG_SERVICES, signedToken, err, "auth[Login]: rawToken.SignedString failed to create jwt token")
		panic(err)
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(constant.ACCESS_TOKEN, signedToken, 3600*24*config.GetInstance().JWT.ExpireTime, "", "", true, true)

	return signedToken
}
