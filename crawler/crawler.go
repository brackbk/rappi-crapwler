package crawler

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"gitlab.com/eiprice/spiders/rappi/repositories"
)

type Crawler struct {
	Db        *gorm.DB
	Lat       string
	Lng       string
	Segment   string
	StoreName string
	Category  string
	Scan      string
	// Category    int
	// SubCategory int
}

func (craw *Crawler) Start() {

	typeStoreRepo := repositories.TypestoreRepositoryDb{craw.Db}
	typeStoreCraw := TypeStoreCrawlerInit{typeStoreRepo, craw.Lat, craw.Lng, craw.Scan}

	fmt.Println("request TypeStores from rappi")
	typestore, err := typeStoreCraw.GetData(craw.Segment)

	storeRepo := repositories.StoreRepositoryDb{craw.Db}
	storeCraw := StoreCrawlerInit{storeRepo, craw.Lat, craw.Lng, craw.Scan}

	categoryRepo := repositories.CategoryRepositoryDb{craw.Db}
	categoryCraw := CategoryCrawlerInit{categoryRepo, craw.Lat, craw.Lng, craw.Category, craw.Scan}

	productRepo := repositories.ProductRepositoryDb{craw.Db}
	productCraw := ProductCrawlerInit{productRepo, craw.Lat, craw.Lng, craw.Scan}

	if err != nil {
		log.Fatalf("Error Get TypeStore Data: %v", err)
	}

	for _, item := range typestore {

		if len(item.SubOptions) > 0 {

			fmt.Println("Get Stores from: ", item.TypeStore)
			stores, err := storeCraw.GetData(item, craw.StoreName)
			if err != nil {
				log.Fatalf("Error Get Store Data : %v", err)
			}

			categories, _ := categoryCraw.GetData(stores)

			for _, category := range categories {
				productCraw.GetData(category, item)
			}

		}
	}
}
