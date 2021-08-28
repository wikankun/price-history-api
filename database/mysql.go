package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type Config struct {
	Host     string
	User     string
	Password string
	Database string
}

var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.User, config.Password, config.Host, config.Database)
	return connectionString
}

//Connector variable used for CRUD operation's
var Connector *gorm.DB

//Connect creates MySQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connected to Database")
	return nil
}
