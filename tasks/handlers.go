package tasks

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/hibiken/asynq"
	"github.com/wikankun/price-history-api/database"
	"github.com/wikankun/price-history-api/entity"
)

func HandlePriceUpdate(ctx context.Context, t *asynq.Task) error {
	type request struct {
		Url string `json:"url"`
	}

	type response struct {
		Name  string  `josn:"name"`
		Price float32 `json:"price"`
	}

	var p PriceTaskPayload
	err := json.Unmarshal(t.Payload(), &p)
	if err != nil {
		return err
	}

	// get item from database
	var item entity.Item
	tx := database.Connector.First(&item, p.ItemID)
	if tx.Error != nil {
		return nil
	}

	// create request
	req := request{
		Url: item.Url,
	}

	jsonReq, err := json.Marshal(req)
	if err != nil {
		return err
	}

	// post to scrapper api
	resp, err := http.Post(
		os.Getenv("SCRAPPER_HOST"),
		"application/json",
		bytes.NewBuffer(jsonReq),
	)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// crate new Price History
	var res response
	json.Unmarshal(bodyBytes, &res)

	price := entity.PriceHistory{
		Item_ID: p.ItemID,
		Price:   uint(res.Price),
	}

	database.Connector.Create(&price)

	return nil
}
