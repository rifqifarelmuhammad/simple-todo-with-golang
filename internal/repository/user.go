package repository

import (
	"github.com/google/uuid"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/constant"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/database"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/models"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/log"
)

func CreateUser(email string, hashedPassword []byte) *models.User {
	user := &models.User{
		UID:      (uuid.New()).String(),
		Email:    email,
		Password: string(hashedPassword),
	}

	result := database.GetInstance().Create(user)
	if result.Error != nil {
		log.Error(constant.TAG_REPOSITORY, result, result.Error, "user[CreateUser]: Error query db on database.GetInstance().Create(args)")
		panic(result.Error)
	}

	return user
}

func FindUserByEmail(email string) *models.User {
	user := &models.User{}
	result := database.GetInstance().Find(user, models.User{Email: email})
	if result.Error != nil {
		log.Error(constant.TAG_REPOSITORY, result, result.Error, "user[FindUserByEmail]: Error query db on database.GetInstance().Find(args)")
		panic(result.Error)
	}

	return user
}

func FindUserByUid(uid string) *models.User {
	user := &models.User{}
	result := database.GetInstance().Find(user, models.User{UID: uid})
	if result.Error != nil {
		log.Error(constant.TAG_REPOSITORY, result, result.Error, "user[FindUserByUid]: Error query db on database.GetInstance().Find(args)")
		panic(result.Error)
	}

	return user
}

func UpdatePassword(user *models.User, hashedNewPassword []byte) {
	user.Password = string(hashedNewPassword)
	result := database.GetInstance().Save(user)
	if result.Error != nil {
		log.Error(constant.TAG_REPOSITORY, result, result.Error, "user[UpdatePassword]: Error query db on database.GetInstance().Save(args)")
		panic(result.Error)
	}
}
