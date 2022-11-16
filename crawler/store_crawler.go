package crawler

import (
	"fmt"

	"gitlab.com/eiprice/spiders/rappi/domain"
	"gitlab.com/eiprice/spiders/rappi/repositories"
	"gitlab.com/eiprice/spiders/rappi/utils"
)

type StoreCrawler interface {
	GetData(StoreType string, StoreTypeName string, StoreName string) ([]*domain.Store, interface{}, error)
	GetDataById(StoreType string, StoreTypeName string, StoreName string, StoreId string) (*domain.Store, interface{}, error)
}

type StoreCrawlerInit struct {
	StoreRepository repositories.StoreRepository
	Lat             string
	Lng             string
	List            string
}

func (craw *StoreCrawlerInit) GetData(typeStore *domain.TypeStore, StoreName string) ([]*domain.Store, error) {

	var storesData []*domain.Store

	var dados interface{}
	var headers []utils.Header
	token, err := Auth()

	if err != nil {
		return nil, fmt.Errorf("Error to get Token from Rappi ")
	}

	headers = append(headers, utils.Header{
		Key:   "Authorization",
		Value: "Bearer " + token,
	})

	payload := `{"state":{"lat":"` + craw.Lat + `","lng":"` + craw.Lng + `","parent_store_type":"` + typeStore.TypeStore + `"},"limit":20000,"offset":0,"context":"cpgs_landing"}`
	url := "https://services.rappi.com.br/api/dynamic/context/content"
	fmt.Println(url)
	fmt.Println(payload)
	dados, err = utils.Request("POST", url, payload, headers)

	if err != nil || dados == nil || dados == "" || dados.(map[string]interface{})["data"].(map[string]interface{})["components"] == nil {
		return nil, err
	}
	fmt.Println("passou 1")
	for _, components := range dados.(map[string]interface{})["data"].(map[string]interface{})["components"].([]interface{}) {
		component := components.(map[string]interface{})
		if component["name"].(string) == "store_groups" {

			for _, stores := range component["resource"].(map[string]interface{})["stores"].([]interface{}) {
				store := stores.(map[string]interface{})
				fmt.Println(store["description"].(string), store["name"].(string))
				if StoreName != "" {
					if store["description"].(string) == StoreName || store["name"].(string) == StoreName {
						obj, _ := craw.GetDataById(typeStore, StoreName, store)
						if obj != nil {
							storesData = append(storesData, obj)
						}
					}
				} else {
					obj, _ := craw.GetDataById(typeStore, StoreName, store)
					if obj != nil {
						storesData = append(storesData, obj)
					}

				}
			}
		}

	}

	craw.StoreRepository.InsertMany(storesData)

	return storesData, nil

}

func (craw *StoreCrawlerInit) GetDataById(typeStore *domain.TypeStore, StoreName string, Component map[string]interface{}) (*domain.Store, error) {

	var store *domain.Store

	if Component != nil {

		fmt.Println("get store : ", utils.StringNotNull(Component["name"]))
		store, _ = domain.NewStore(
			utils.StringNotNull(Component["name"]),
			typeStore.TypeStore,
			typeStore.Name,
			int(Component["zone_id"].(float64)),
			int(Component["store_id"].(float64)),
			utils.StringNotNull(Component["description"]),
			Component["image"].(string),
			" ",
			" ",
			" ",
			utils.FloatNotNull(Component["delivery_price"]),
			utils.FloatNotNull(Component["eta_value"]),
			0,
			" ",
			fmt.Sprint(utils.FloatNotNull(Component["lat"])),
			fmt.Sprint(utils.FloatNotNull(Component["lng"])),
			craw.List,
		)
	}

	return store, nil

}
