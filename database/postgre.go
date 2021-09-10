package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	User     string
	Password string
	Database string
	Port     string
}

var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.Host, config.User, config.Password, config.Database, config.Port)
	return connectionString
}

//Connector variable used for CRUD operation's
var Connector *gorm.DB

//Connect creates MySQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}
	log.Println("Connected to Database")
	return nil
}
