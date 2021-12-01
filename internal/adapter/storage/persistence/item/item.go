package item

import "github.com/LibenHailu/addisber_scraper/internal/constant/model"

// ItemPersistence contains a list of functions for database operation for table items
type ItemPersistence interface {
	UpdateItem(item *model.Item) error
	SearchItem(seatch string) ([]model.Item, error)
}
