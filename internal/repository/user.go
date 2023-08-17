package repository

import (
	"github.com/google/uuid"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/constant"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/database"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/models"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/log"
)

func CreateUser(email string, hashedPassword []byte) error {
	user := models.User{
		UID:      uuid.New(),
		Email:    email,
		Password: string(hashedPassword),
	}

	result := database.GetInstance().Create(user)
	if result.Error != nil {
		log.Error(constant.TAG_REPOSITORY, result, result.Error, "user[CreateUser]: Error query db on database.GetInstance().Create")
	}

	return result.Error
}

func FindUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	result := database.GetInstance().Find(user, models.User{Email: email})
	if result.Error != nil {
		log.Error(constant.TAG_REPOSITORY, result, result.Error, "user[FindUserByEmail]: Error query db on database.GetInstance().Find")
	}

	return user, result.Error
}

func FindUserByUid(uid uuid.UUID) (*models.User, error) {
	user := &models.User{}
	result := database.GetInstance().Find(user, models.User{UID: uid})
	if result.Error != nil {
		log.Error(constant.TAG_REPOSITORY, result, result.Error, "user[FindUserByUid]: Error query db on database.GetInstance().Find")
	}

	return user, result.Error
}
