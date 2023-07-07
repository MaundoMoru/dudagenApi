package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// load env file
func GoDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}

// connect to databse and return it as object
func DbConnect() (db *gorm.DB) {
	// pass db variables into variables
	host := GoDotEnvVariable("HOST")
	port := GoDotEnvVariable("PORT")
	user := GoDotEnvVariable("USER")
	password := GoDotEnvVariable("PASSWORD")
	dbname := GoDotEnvVariable("DBNAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	println("database url is " + psqlInfo)

	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")
	return db
}
