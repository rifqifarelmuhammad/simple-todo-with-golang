package database

import (
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/constant"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/models"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/log"
)

func Migrate() {
	var err error

	err = pool.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(constant.TAG_DATABASE, err, "Failed to migrate user model")
	}

	err = pool.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatal(constant.TAG_DATABASE, err, "Failed to migrate todo model")
	}

	log.Print(constant.TAG_DATABASE, "Database Migration Completed")
}
