package database

import (
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/constant"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/models"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/log"
)

func Migrate() {
	err := pool.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
		log.Fatal(constant.TAG_DATABASE, err, "Failed to migrate")
	}

	log.Print(constant.TAG_DATABASE, "Database Migration Completed")
}
