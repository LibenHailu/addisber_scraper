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
	is.c.OnXML(`//div[@class="shop-loop-after clearfix"]/nav[@class="woocommerce-pagination"]/ul[@class="page-numbers"]/li/a[@class="next page-numbers"]`, func(e *colly.XMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	is.c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Set error handler
	is.c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	is.c.OnScraped(func(r *colly.Response) {
		is.c.Visit("https://www.addisber.com/shop/")
	})

	is.c.Visit("https://www.addisber.com/shop/")
}
