package crawler

import (
	"fmt"
	"strconv"

	"gitlab.com/eiprice/spiders/rappi/domain"
	"gitlab.com/eiprice/spiders/rappi/repositories"
	"gitlab.com/eiprice/spiders/rappi/utils"
)

type ProductCrawler interface {
	GetData(categories []*domain.Category, typeStore *domain.TypeStore)
	GetById(category *domain.Category, typeStore *domain.TypeStore) ([]*domain.Product, error)
}

type ProductCrawlerInit struct {
	ProductRepository repositories.ProductRepository
	Lat               string
	Lng               string
	List              string
}

func (craw *ProductCrawlerInit) GetData(categories []*domain.Category, typeStore *domain.TypeStore) {

	for _, item := range categories {
		fmt.Println("Get product from Category: ", item.Name)
		product, err := craw.GetById(item, typeStore)
		craw.ProductRepository.InsertMany(product)

		if err != nil {
			fmt.Println("Error to get Products  from Category:", item.Name)
		}

		fmt.Println("Total products inserted from Category "+item.Name+": ", len(product))
	}

}

func (craw *ProductCrawlerInit) GetById(category *domain.Category, typeStore *domain.TypeStore) ([]*domain.Product, error) {

	var dados interface{}
	var product []*domain.Product
	var headers []utils.Header
	token, err := Auth()

	if err != nil {
		return nil, fmt.Errorf("Error to get Token from Rappi ")
	}

	headers = append(headers, utils.Header{
		Key:   "Authorization",
		Value: "Bearer " + token,
	})

	//url := "https://services.rappi.com.br/api/subcorridor_sections/products?subcorridor_id=" + strconv.Itoa(category.SubCategoryID) + "&store_id=" + strconv.Itoa(category.StoreID) + "&include_stock_out=true&limit=999"
	url := "https://services.rappi.com.br/api/dynamic/context/content"
	fmt.Println(url)

	payload := `{"state":{"aisle_id":"` + strconv.Itoa(category.SubCategoryID) + `"},"limit":9999,"offset":0,"context":"aisle_detail","stores":[` + strconv.Itoa(category.StoreID) + `]}`

	fmt.Println(payload)

	dados, err = utils.Request("POST", url, payload, headers)

	if err != nil || dados == nil {
		return nil, err
	}

	for _, item := range dados.(map[string]interface{})["data"].(map[string]interface{})["components"].([]interface{}) {
		for _, subItem := range item.(map[string]interface{})["resource"].(map[string]interface{})["products"].([]interface{}) {

			if subItem.(map[string]interface{})["product_id"] != nil {
				productID, _ := strconv.Atoi(subItem.(map[string]interface{})["product_id"].(string))

				obj, _ := domain.NewProduct(
					utils.FloatNotNull(subItem.(map[string]interface{})["discout"]),
					productID,
					utils.Ean(utils.StringNotNull(subItem.(map[string]interface{})["ean"])),
					utils.BoolNotNull(subItem.(map[string]interface{})["in_stock"]),
					utils.StringNotNull(subItem.(map[string]interface{})["trademark"]),
					category.CategoryID,
					category.Name,
					category.SubCategoryID,
					category.SubName,
					utils.StringNotNull(subItem.(map[string]interface{})["description"]),
					`https://images.rappi.com.br/products/`+utils.StringNotNull(subItem.(map[string]interface{})["image"]),
					utils.BoolNotNull(subItem.(map[string]interface{})["is_available"]),
					utils.StringNotNull(subItem.(map[string]interface{})["name"]),
					utils.FloatNotNull(subItem.(map[string]interface{})["price"]),
					utils.FloatNotNull(subItem.(map[string]interface{})["min_quantity_in_grams"]),
					utils.StringNotNull(subItem.(map[string]interface{})["pum"]),
					utils.IntNotNull(subItem.(map[string]interface{})["quantity"]),
					utils.FloatNotNull(subItem.(map[string]interface{})["real_price"]),
					category.StoreID,
					category.StoreName,
					utils.FloatNotNull(subItem.(map[string]interface{})["real_balance_price"]),
					utils.StringNotNull(subItem.(map[string]interface{})["retail_id"]),
					utils.StringNotNull(subItem.(map[string]interface{})["sale_type"]),
					utils.StringNotNull(subItem.(map[string]interface{})["store_type"]),
					utils.StringNotNull(subItem.(map[string]interface{})["unit_type"]),
					typeStore.TypeStore,
					typeStore.Name,
					craw.List,
					`https://www.rappi.com.br/produto/`+strconv.Itoa(category.StoreID)+`_`+strconv.Itoa(productID),
				)
				product = append(product, obj)
			}
		}
	}

	return product, nil
}
