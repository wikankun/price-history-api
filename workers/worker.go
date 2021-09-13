package main

import (
	"log"
	"os"

	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
	"github.com/wikankun/price-history-api/database"
	"github.com/wikankun/price-history-api/tasks"
)

func main() {
	godotenv.Load()

	config :=
		database.Config{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("USERNAME"),
			Password: os.Getenv("PASSWORD"),
			Database: os.Getenv("DATABASE"),
			Port:     os.Getenv("DB_PORT"),
		}

	initDB(config)

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: os.Getenv("REDIS_HOST")},
		asynq.Config{Concurrency: 1},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(
		tasks.TypePriceUpdate,   // task type
		tasks.HandlePriceUpdate, // handler function
	)

	err := srv.Run(mux)
	if err != nil {
		log.Fatal(err)
	}
}

func initDB(config database.Config) {
	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
}
