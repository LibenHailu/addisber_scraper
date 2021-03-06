package item

import (
	"github.com/LibenHailu/addisber_scraper/internal/constant/model"
)

// UpdateItem applies update item bussness logic
func (s *service) UpdateItem(item *model.Item) error {
	err := s.itemPersist.UpdateItem(item)
	return err
}

// SeatchItem applies search item business logic
func (s *service) SearchItem(search string) ([]model.Item, error) {
	return s.itemPersist.SearchItem(search)
}
