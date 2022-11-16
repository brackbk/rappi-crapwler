package crawler

import (
	"fmt"
	"reflect"
	"strconv"

	"gitlab.com/eiprice/spiders/rappi/domain"
	"gitlab.com/eiprice/spiders/rappi/repositories"
	"gitlab.com/eiprice/spiders/rappi/utils"
)

type CategoryCrawler interface {
	GetData(stores []*domain.Store) ([][]*domain.Category, error)
	GetById(stores *domain.Store) ([]*domain.Category, error)
	GetByIdCorridor(stores *domain.Store) ([]*domain.Category, error)
}
type CategoryCrawlerInit struct {
	CategoryRepository repositories.CategoryRepository
	Lat                string
	Lng                string
	Category           string
	List               string
}

func (craw *CategoryCrawlerInit) GetData(stores []*domain.Store) ([][]*domain.Category, error) {
	var categories [][]*domain.Category

	for _, item := range stores {

		fmt.Println("Get Categories from store: ", item.Name)

		category, err := craw.GetById(item)
		if err != nil {
			fmt.Println("Error to get Categories from Store:", item.Name)
		}
		craw.CategoryRepository.InsertMany(category)
		fmt.Println("Total Categories inserted from Store "+item.Name+": ", len(category))
		categories = append(categories, category)
	}

	return categories, nil
}

func (craw *CategoryCrawlerInit) GetByIdCorridor(stores *domain.Store) ([]*domain.Category, error) {

	var dados interface{}
	var headers []utils.Header
	token, err := Auth()
	fmt.Println(token)
	var category []*domain.Category
	headers = append(headers, utils.Header{
		Key:   "Authorization",
		Value: "Bearer " + token,
	})

	url := "https://services.rappi.com.br/windu/corridors/store/" + strconv.Itoa(stores.StoreID)
	fmt.Println(url)
	dados, err = utils.Request("GET", url, "", headers)

	if err != nil || dados == nil {
		return nil, err
	}

	for _, item := range dados.(map[string]interface{})["corridors"].([]interface{}) {

		obj, _ := domain.NewCategory(
			int(stores.StoreID),
			int(item.(map[string]interface{})["corridor_mapper_id"].(float64)),
			int(item.(map[string]interface{})["corridor_id"].(float64)),
			item.(map[string]interface{})["name"].(string),
			item.(map[string]interface{})["name"].(string),
			stores.Name,
			craw.List,
		)
		if craw.Category != "" {
			if item.(map[string]interface{})["name"].(string) == craw.Category {
				category = append(category, obj)
			}
		} else {
			category = append(category, obj)
		}

	}

	return category, nil
}

func (craw *CategoryCrawlerInit) GetById(stores *domain.Store) ([]*domain.Category, error) {

	var dados interface{}
	var headers []utils.Header
	token, err := Auth()
	fmt.Println(token)
	var category []*domain.Category
	headers = append(headers, utils.Header{
		Key:   "Authorization",
		Value: "Bearer " + token,
	})

	url := "https://services.rappi.com.br/windu/corridors/sub_corridors/store/" + strconv.Itoa(stores.StoreID)
	fmt.Println(url)
	dados, err = utils.Request("GET", url, "", headers)

	if err != nil || dados == nil {
		return nil, err
	}

	if reflect.ValueOf(dados).Kind() == reflect.Map {
		category, err = craw.GetByIdCorridor(stores)
		return category, err
	} else {
		for _, item := range dados.([]interface{}) {

			for _, subItem := range item.(map[string]interface{})["sub_corridors"].([]interface{}) {

				obj, _ := domain.NewCategory(
					int(subItem.(map[string]interface{})["store_id"].(float64)),
					int(item.(map[string]interface{})["corridor_mapper_id"].(float64)),
					int(subItem.(map[string]interface{})["subcorridor_id"].(float64)),
					item.(map[string]interface{})["name"].(string),
					subItem.(map[string]interface{})["name"].(string),
					stores.Name,
					craw.List,
				)
				if craw.Category != "" {
					if item.(map[string]interface{})["name"].(string) == craw.Category {
						category = append(category, obj)
					}
				} else {
					category = append(category, obj)
				}

			}
		}
	}

	return category, nil
}
