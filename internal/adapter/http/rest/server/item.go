package server

import (
	"net/http"

	"github.com/LibenHailu/addisber_scraper/internal/module/item"
	"github.com/gin-gonic/gin"
)

// ItemHandler contains a function of hadlers for the domain item
type ItemHandler interface {
	SearchItem(c *gin.Context)
}

// itemHandler defines all the necessary things for item handler
type itemHandler struct {
	itemUsecase item.Usecase
}

// ItemInit initialize an item handler for the domain item
func ItemInit(itemUsecase item.Usecase) ItemHandler {
	return &itemHandler{
		itemUsecase,
	}
}

// SearchItem searches for an item
// GET /api/v1/search
func (ih itemHandler) SearchItem(c *gin.Context) {
	items, _ := ih.itemUsecase.SearchItem(c.Query("q"))

	c.JSON(http.StatusOK, gin.H{"items": items})
}
