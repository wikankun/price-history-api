package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/wikankun/price-history-api/database"
	"github.com/wikankun/price-history-api/entity"
)

//GetPriceHistoryByID returns price history with specific ID
func GetPriceHistoryByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["item_id"])

	var prices []entity.PriceHistory
	database.Connector.Find(&prices, entity.PriceHistory{Item_ID: key})
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
	type request struct {
		Url string `json:"url"`
	}

	type response struct {
		Name  string  `josn:"name"`
		Price float32 `json:"price"`
	}

	// get id
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["item_id"])

	// get item from database
	var item entity.Item
	database.Connector.First(&item, key)

	// create request
	req := request{
		Url: item.Url,
	}

	jsonReq, err := json.Marshal(req)

	// post to scrapper api
	resp, err := http.Post(os.Getenv("SCRAPPER_HOST"), "application/json", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// crate new Price History
	var res response
	json.Unmarshal(bodyBytes, &res)

	price := entity.PriceHistory{
		Item_ID:   item.ID,
		Price:     uint(res.Price),
		UpdatedAt: time.Now(),
	}

	database.Connector.Create(&price)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(price)
}
