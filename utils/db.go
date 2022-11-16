package utils

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"gitlab.com/eiprice/spiders/rappi/domain"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}

func ConnectDB(drop string) *gorm.DB {
	var dsn string
	var db *gorm.DB
	var err error

	dsn = os.Getenv("dsn")
	db, err = gorm.Open(os.Getenv("dbType"), dsn)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	if os.Getenv("debug") == "true" {
		db.LogMode(true)
	}

	if drop == "all" {
		log.Printf("Cleaning Database...")
		db.Exec("TRUNCATE rappi.products;")
		db.Exec("TRUNCATE rappi.categories;")
		db.Exec("TRUNCATE rappi.stores;")
		db.Exec("TRUNCATE rappi.type_stores;")
		log.Printf("Database is clean")
	} else {
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return os.Getenv("schema") + "." + defaultTableName
		}

		if os.Getenv("AutoMigrateDb") == "true" {
			db.AutoMigrate(&domain.TypeStore{}, &domain.Store{}, &domain.Product{}, &domain.Category{})
			os.Setenv("AutoMigrateDb", "false")
		}
	}

	return db
}
