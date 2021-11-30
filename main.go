package x

// import (
// 	"encoding/csv"
// 	"log"
// 	"os"

// 	"github.com/gocolly/colly"
// )

// func main() {
// 	fName := "addisber_items.csv"
// 	file, err := os.Create(fName)
// 	if err != nil {
// 		log.Fatalf("Cannot create file %q: %s\n", fName, err)
// 		return
// 	}
// 	defer file.Close()
// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()
// 	// Write CSV header
// 	writer.Write([]string{"ID", "Title", "Description", "Price", "URL", "Image URL"})

// 	// Instantiate default collector
// 	c := colly.NewCollector(
// 	// Allow requests only to store.xkcd.com
// 	// colly.AllowedDomains("addisber.com"),
// 	)

// 	// Extract product details
// 	c.OnHTML(".product-inner", func(e *colly.HTMLElement) {
// 		writer.Write([]string{
// 			e.ChildText(".woocommerce-loop-product__title"),
// 			e.ChildText(".woocommerce-loop-product__title"),
// 			e.ChildText(".woocommerce-Price-amount bdi"),
// 			e.Request.AbsoluteURL(e.ChildAttr("a", "href")),
// 			e.ChildAttr("img", "src"),
// 		})

// 	})

// 	// Find and visit next page links
// 	c.OnHTML(`.page-numbers a[href]`, func(e *colly.HTMLElement) {
// 		e.Request.Visit(e.Attr("href"))
// 	})

// 	c.Visit("https://www.addisber.com/shop/")

// 	log.Printf("Scraping finished, check file %q for results\n", fName)

// 	// Display collector's statistics
// 	log.Println(c)

// }
