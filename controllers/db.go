package controllers

import (
	"fmt"
	"os"

	"github.com/ervinismu/kejarmimpi/models"
	"github.com/jinzhu/gorm"
	//For connect postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//InitDb is for connecting to database
func InitDb() *gorm.DB {
	var err error
	dbhost := os.Getenv("DATABASE_URL")
	if dbhost == "" {
		dbhost = "host=127.0.0.1 user=postgres dbname=kejarmimpi sslmode=disable password=postgresreza"
	}
	var post models.Post
	var user models.User
	db, err := gorm.Open("postgres", dbhost)
	db.AutoMigrate(post, user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection Success!")
	}
	return db
}
