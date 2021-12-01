package initiator

import (
	"log"
	"os"

	scraper "github.com/LibenHailu/addisber_scraper/internal/adapter/scraper/colly"
	itemPersist "github.com/LibenHailu/addisber_scraper/internal/adapter/storage/persistence/item"
	"github.com/LibenHailu/addisber_scraper/internal/constant/model"
	"github.com/LibenHailu/addisber_scraper/internal/module/item"
	"github.com/gocolly/colly"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initialize() {
	dsn := "postgres://postgres:admin@localhost:5432/scrap?sslmode=disable"
	// TODO: use connection pool
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Printf("Error when Opening database connection: %v", err)
		os.Exit(1)
	}

	// conn.AutoMigrate migrates gorm models
	conn.AutoMigrate(&model.Item{})

	itemPersistence := itemPersist.ItemInit(conn)

	itemUsecase := item.Initialize(itemPersistence)

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowURLRevisit(),
		// colly.Debugger(&debug.LogDebugger{}),
		// colly.Async(true),
	)
	// c.Wait()

	scrapHandler := scraper.ItemInit(c, itemUsecase)

	scrapHandler.UpdateItems()

}
