package colly

import (
	"fmt"

	"github.com/LibenHailu/addisber_scraper/internal/constant/model"
	"github.com/LibenHailu/addisber_scraper/internal/module/item"
	"github.com/gocolly/colly"
)

// ItemScraper contains  a functions of scrapers for domain item
type ItemScraper interface {
	UpdateItems()
}

// itemScraper defiens all the necessary things for scrap handling
type itemScraper struct {
	c           *colly.Collector
	itemUsecase item.Usecase
}

// ItemInit initializes an item scraper for the domain item
func ItemInit(c *colly.Collector, itemUsecase item.Usecase) ItemScraper {
	return &itemScraper{
		c,
		itemUsecase,
	}
}

func (is *itemScraper) UpdateItems() {
	// Extract product details
	is.c.OnHTML(".product-inner", func(e *colly.HTMLElement) {
		item := model.Item{
			AddisberID: e.ChildText(".woocommerce-loop-product__title"),
			Title:      e.ChildText(".woocommerce-loop-product__title"),
			Price:      e.ChildText(".woocommerce-Price-amount bdi"),
			URL:        e.Request.AbsoluteURL(e.ChildAttr("a", "href")),
			PictureURL: e.ChildAttr("img", "src"),
		}

		is.itemUsecase.UpdateItem(&item)

	})

	// Find and visit next page links
	is.c.OnHTML(`.next a[href]`, func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	is.c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	for {
		is.c.Visit("https://store.xkcd.com/collections/everything")
	}
}
