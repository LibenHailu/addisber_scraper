package routing

import (
	"github.com/LibenHailu/addisber_scraper/internal/adapter/http/rest/server"
	"github.com/gin-gonic/gin"
)

// ItemRoutes maps route string to its handler for the domain item
func ItemRoutes(v1 *gin.RouterGroup, itemHandler server.ItemHandler) {
	v1.GET("/search", itemHandler.SearchItem)
}
