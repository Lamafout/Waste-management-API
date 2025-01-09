package main

import (
	"log"
	"os"
	"waste_management/config"
	"waste_management/model"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Println("Error while reading dotenv: ", err.Error())
	}

	conf := &config.MongoDbConfig{
		Uri: os.Getenv("MONGODB_URI"),
		Database: os.Getenv("MONGODB_DATABASE"),
	}

	connection := model.NewMongoDBConnection(conf)

	log.Println(connection)
}