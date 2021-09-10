package migrations

import (
	"github.com/wikankun/price-history-api/database"
	"github.com/wikankun/price-history-api/entity"
)

func Migrate() {
	database.Connector.AutoMigrate(&entity.Item{}, &entity.PriceHistory{})
}
