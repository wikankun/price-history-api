package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const (
	// TypePriceUpdate is task type for Price Update
	TypePriceUpdate = "price:update"
)

// PriceTaskPayload is a payload for Price Task
type PriceTaskPayload struct {
	ItemID int
}

// NewPriceUpdateTask is to create task for Price Update
func NewPriceUpdateTask(itemID int) (*asynq.Task, error) {
	payload, err := json.Marshal(PriceTaskPayload{ItemID: itemID})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypePriceUpdate, payload), nil
}
