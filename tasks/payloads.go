package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const (
	TypePriceUpdate = "price:update"
)

type PriceTaskPayload struct {
	ItemID int
}

func NewPriceUpdateTask(item_id int) (*asynq.Task, error) {
	payload, err := json.Marshal(PriceTaskPayload{ItemID: item_id})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypePriceUpdate, payload), nil
}
