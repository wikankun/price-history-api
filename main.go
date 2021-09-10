package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/wikankun/price-history-api/controllers"
	"github.com/wikankun/price-history-api/database"
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
	port := os.Getenv("PORT")

	initDB(config)
	// migrations.Migrate()
	log.Printf("Starting HTTP Server on port %s", port)
	router := mux.NewRouter().StrictSlash(true)
	initHandlers(router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT"},
	})

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func initHandlers(router *mux.Router) {
	router.HandleFunc("/item", controllers.CreateItem).Methods("POST")
	router.HandleFunc("/item", controllers.GetAllItem).Methods("GET")
	router.HandleFunc("/item/{id}", controllers.GetItemByID).Methods("GET")
	router.HandleFunc("/item/{id}", controllers.UpdateItemByID).Methods("PUT")
	// router.HandleFunc("/delete/{id}", controllers.DeleteItemByID).Methods("DELETE")
	router.HandleFunc("/price/{item_id}", controllers.GetPriceHistoryByID).Methods("GET")
	router.HandleFunc("/price", controllers.CreatePriceHistory).Methods("POST")
}

func initDB(config database.Config) {
	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
}
