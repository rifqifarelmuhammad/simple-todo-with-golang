package database

import (
	"fmt"

	"github.com/rifqifarelmuhammad/simple-todo-with-golang/config"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/constant"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var pool *gorm.DB

func GetInstance() *gorm.DB {
	return pool
}

func Init() {
	ConnectToDatabase()
	Migrate()
}

func ConnectToDatabase() {
	dbConfig := config.GetInstance().Database
	dbStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.SSLMode,
	)

	var err error
	pool, err = gorm.Open(postgres.Open(dbStr), &gorm.Config{})
	if err != nil {
		log.Fatal(constant.TAG_DATABASE, err, "Failed to connect database")
	}

	log.Print(constant.TAG_DATABASE, "Connected to Database")
}

func CloseConnection() {
	db, _ := pool.DB()
	err := db.Close()
	if err != nil {
		log.Fatal(constant.TAG_DATABASE, err, "Can not close database connection")
	}
}
