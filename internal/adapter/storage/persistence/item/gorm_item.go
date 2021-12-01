package item

import (
	"errors"

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

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = gp.conn.Create(&item).Error
		// fmt.Println(err)
		return err
	}

	err = gp.conn.Model(&model.Item{}).Where("id = ?", foundItem.ID).Updates(*item).Error
	return err
}

// SearchItem searches item on the database
func (gp *itemGormPersistence) SearchItem(search string) ([]model.Item, error) {
	// TODO: make this search order by best paramerter like rating
	var items []model.Item

	err := gp.conn.Raw("SELECT * FROM items WHERE title LIKE '%" + search + "%';").Scan(&items).Error

	return items, err
}
