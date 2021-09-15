package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/wikankun/price-history-api/database"
	"github.com/wikankun/price-history-api/entity"
	"github.com/wikankun/price-history-api/helpers"
)

//GetAllItem get all item data
func GetAllItem(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	var items []entity.Item
	scopes := database.Connector.Scopes(helpers.Paginate(r))
	if name == "" {
		scopes.Order("id").Find(&items)
	} else {
		query := fmt.Sprintf("%%%s%%", strings.ToLower(name))
		scopes.Order("id").Where("LOWER(name) LIKE ?", query).Find(&items)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

//GetItemByID returns item with specific ID
func GetItemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var item entity.Item
	database.Connector.First(&item, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

//CreateItem creates item
func CreateItem(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var item entity.Item
	json.Unmarshal(requestBody, &item)

	database.Connector.Create(&item)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

//UpdateItemByID updates item with respective ID
func UpdateItemByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var item entity.Item
	json.Unmarshal(requestBody, &item)

	database.Connector.Save(&item)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

//DeleteItemByID delete's item with specific ID
func DeleteItemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var item entity.Item
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&item)
	w.WriteHeader(http.StatusNoContent)
}
