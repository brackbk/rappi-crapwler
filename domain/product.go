package domain

import (
	"time"
)

type Product struct {
	Base
	Discount              float64 `json:"title" gorm:"type:varchar(255)"`
	Product_id            int     `json:"product_id"`
	Ean                   string  `json:"ean" gorm:"type:varchar(255)" `
	In_stock              bool    `json:"In_stock`
	Brand_name            string  `json:"brand_name" gorm:"type:varchar(255)"`
	Corridor_id           int     `json:"corridor_id"`
	Corridor_name         string  `json:"corridor_name" gorm:"type:varchar(255)"`
	Sub_corridor_id       int     `json:"sub_corridor_id"`
	Sub_corridor_name     string  `json:"sub_corridor_name" gorm:"type:varchar(255)"`
	Description           string  `json:"description" gorm:"type:text"`
	Image                 string  `json:"image" gorm:"type:text"`
	Is_available          bool    `json:"is_available"`
	Name                  string  `json:"name" gorm:"type:varchar(255)"`
	Price                 float64 `json:"price"`
	Min_quantity_in_grams float64 `json:"min_quantity_in_grams"`
	Pum                   string  `json:"pum"`
	Quantity              int     `json:"quantity"`
	Real_price            float64 `json:"real_price" gorm:"type:varchar(255)"`
	Store_id              int     `json:"store_id"`
	Store_name            string  `json:"store_name" gorm:"type:varchar(255)"`
	Real_balance_price    float64 `json:"real_balance_price"`
	Retail_id             string  `json:"retail_id" gorm:"type:varchar(255)"`
	Sale_type             string  `json:"sale_type" gorm:"type:varchar(255)"`
	Store_type            string  `json:"store_type" gorm:"type:varchar(255)"`
	Unit_type             string  `json:"unit_type" gorm:"type:varchar(255)"`
	TypeStore             string  `json:"typestore" gorm:"type:varchar(255)"`
	TypeStoreName         string  `json:"typestorename" gorm:"type:varchar(255)"`
	Scan                  string  `json:"scan" gorm:"type:varchar(255)"`
	Url                   string  `json:"url" gorm:"type:varchar(255)"`
}

func NewProduct(
	Discount float64,
	Product_id int,
	Ean string,
	In_stock bool,
	Brand_name string,
	Corridor_id int,
	Corridor_name string,
	Sub_corridor_id int,
	Sub_corridor_name string,
	Description string,
	Image string,
	Is_available bool,
	Name string,
	Price float64,
	Min_quantity_in_grams float64,
	Pum string,
	Quantity int,
	Real_price float64,
	Store_id int,
	Store_name string,
	Real_balance_price float64,
	Retail_id string,
	Sale_type string,
	Store_type string,
	Unit_type string,
	TypeStore string,
	TypeStoreName string,
	Scan string,
	Url string,
) (*Product, error) {

	product := &Product{
		Discount:              Discount,
		Product_id:            Product_id,
		Ean:                   Ean,
		In_stock:              In_stock,
		Brand_name:            Brand_name,
		Corridor_id:           Corridor_id,
		Corridor_name:         Corridor_name,
		Sub_corridor_id:       Sub_corridor_id,
		Sub_corridor_name:     Sub_corridor_name,
		Description:           Description,
		Image:                 Image,
		Is_available:          Is_available,
		Name:                  Name,
		Price:                 Price,
		Min_quantity_in_grams: Min_quantity_in_grams,
		Pum:                   Pum,
		Quantity:              Quantity,
		Real_price:            Real_price,
		Store_id:              Store_id,
		Store_name:            Store_name,
		Real_balance_price:    Real_balance_price,
		Retail_id:             Retail_id,
		Sale_type:             Sale_type,
		Store_type:            Store_type,
		Unit_type:             Unit_type,
		TypeStore:             TypeStore,
		TypeStoreName:         TypeStoreName,
		Scan:                  Scan,
		Url:                   Url,
	}

	//product.ID = uuid.NewV4().String()
	product.CreatedAt = time.Now()

	return product, nil
}
