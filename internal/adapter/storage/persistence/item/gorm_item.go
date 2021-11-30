package item

import (
	"fmt"

	"github.com/LibenHailu/addisber_scraper/internal/constant/model"
	"gorm.io/gorm"
)

// itemGormPersistence defines all the necessary things for the database operation on domian Item
type itemGormPersistence struct {
	conn *gorm.DB
}

// ItemInit creates a new UserPersistence object
func ItemInit(conn *gorm.DB) ItemPersistence {
	return &itemGormPersistence{
		conn,
	}
}

// UpdateItem persists item to the database
func (gp *itemGormPersistence) UpdateItem(item *model.Item) error {
	foundItem := model.Item{}
	err := gp.conn.First(&foundItem, "addisber_id = ?", item.AddisberID).Error

	fmt.Println(err)
	return err
}