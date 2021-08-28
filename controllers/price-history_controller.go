package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wikankun/price-history-api/database"
	"github.com/wikankun/price-history-api/entity"
)

//GetPriceHistoryByID returns item with specific ID
func GetPriceHistoryByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["item_id"])

	var items []entity.PriceHistory
	database.Connector.Find(&items, entity.PriceHistory{Item_ID: key})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

//CreatePriceHistory creates item
func CreatePriceHistory(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var item entity.PriceHistory
	json.Unmarshal(requestBody, &item)

	database.Connector.Create(&item)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
