package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/hibiken/asynq"
	"github.com/wikankun/price-history-api/database"
	"github.com/wikankun/price-history-api/entity"
	"github.com/wikankun/price-history-api/tasks"
)

//GetPriceHistoryByID returns price history with specific ID
func GetPriceHistoryByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["item_id"])

	var prices []entity.PriceHistory
	database.Connector.Order("id").Find(&prices, entity.PriceHistory{ItemID: key})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prices)
}

//CreatePriceHistory creates price history
func CreatePriceHistory(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var price entity.PriceHistory
	json.Unmarshal(requestBody, &price)

	database.Connector.Create(&price)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(price)
}

//UpdatePriceHistory updates price history
func UpdatePriceHistory(w http.ResponseWriter, r *http.Request) {
	// get id
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["item_id"])

	redisOpt, err := asynq.ParseRedisURI(os.Getenv("REDIS_URL"))

	client := asynq.NewClient(redisOpt)
	defer client.Close()

	client.SetDefaultOptions(tasks.TypePriceUpdate, asynq.MaxRetry(2), asynq.Timeout(time.Minute))

	t, err := tasks.NewPriceUpdateTask(key)
	if err != nil {
		log.Fatal(err)
	}

	info, err := client.Enqueue(t)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(info)
}
