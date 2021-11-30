package item

import (
	"github.com/LibenHailu/addisber_scraper/internal/adapter/storage/persistence/item"
	"github.com/LibenHailu/addisber_scraper/internal/constant/model"
)

// Usecase interface contains function of business logic for domain item
type Usecase interface {
	UpdateItem(item *model.Item) error
}

// service defines all necessary service for domain item
type service struct {
	itemPersist item.ItemPersistence
}

// Initialize create a new service object
func Initialize(itemPersist item.ItemPersistence) Usecase {
	return &service{
		itemPersist,
	}
}
