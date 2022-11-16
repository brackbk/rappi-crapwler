package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/joho/sqltocsv"
	_ "github.com/lib/pq"
	"gitlab.com/eiprice/spiders/rappi/crawler"
	"gitlab.com/eiprice/spiders/rappi/utils"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {

	var lat string
	var lng string
	var segment string
	var store string
	var category string
	var scan string
	var drop string
	var file string
	//900020401
	var sc int64

	flag.StringVar(&lat, "lat", "-23.584293", "Set Lat")
	flag.StringVar(&lng, "lng", "-46.674584", "Set Lng")
	flag.StringVar(&segment, "segment", "", "Set Segment")
	flag.StringVar(&store, "store", "", "Set StoreId")
	flag.StringVar(&category, "category", "", "Set Category")
	flag.StringVar(&scan, "scan", "1", "Set Scan")
	flag.StringVar(&drop, "drop", "", "Set Drop")
	flag.StringVar(&file, "file", "", "Set File")
	flag.Parse()

	db := utils.ConnectDB(drop)

	if drop != "all" {
		start := time.Now()
		//Crawler get data from Rappi website

		if file != "" {
			Coords := utils.Read(file)
			fmt.Println(Coords)
			for _, item := range Coords {
				sc = sc + 1

				//if sc > 23 {
				api := crawler.Crawler{db, item.Lat, item.Lng, segment, store, category, fmt.Sprint(sc)}

				log.Printf("Start Crawler from", item.Lat, item.Lng)

				api.Start()

				DB, err := sql.Open("postgres", os.Getenv("dsn"))
				if err != nil {
					panic(err)
				}
				defer DB.Close()
				result, err := DB.Query(`Select * from rappi.products p where p.scan = '` + fmt.Sprint(sc) + `'`)
				if err != nil {
					fmt.Println("error")
					log.Fatal(err)
				}

				csvConverter := sqltocsv.New(result)
				csvConverter.Delimiter = ';'
				csvConverter.WriteFile(`files/` + fmt.Sprint(sc) + `_` + item.Lat + `_` + item.Lng + `.xlsx`)
				//}
			}
		} else {
			api := crawler.Crawler{db, lat, lng, segment, store, category, scan}

			log.Printf("Start Crawler...")

			api.Start()
		}

		elapsed := time.Since(start)
		log.Printf("Finish Crawler Took %s", elapsed)
	}

}
