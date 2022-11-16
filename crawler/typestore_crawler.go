package crawler

import (
	"fmt"

	"gitlab.com/eiprice/spiders/rappi/domain"
	"gitlab.com/eiprice/spiders/rappi/repositories"
	"gitlab.com/eiprice/spiders/rappi/utils"
)

type TypeStoreCrawler interface {
	GetData(segment string) ([]*domain.TypeStore, error)
}
type TypeStoreCrawlerInit struct {
	TypeStoreRepository repositories.TypeStoreRepository
	Lat                 string
	Lng                 string
	List                string
}

func (craw *TypeStoreCrawlerInit) GetData(segment string) ([]*domain.TypeStore, error) {

	var dados interface{}
	var headers []utils.Header
	var typestore []*domain.TypeStore
	token, err := Auth()

	if err != nil {
		return nil, fmt.Errorf("Error to get Token from Rappi ")
	}

	headers = append(headers, utils.Header{
		Key:   "Authorization",
		Value: "Bearer " + token,
	})

	url := "https://services.rappi.com.br/api/sidekick/base-crack/principal?lng=" + craw.Lng + "&lat=" + craw.Lat

	dados, err = utils.Request("GET", url, "", headers)

	if err != nil || dados == nil {
		return nil, err
	}

	for _, item := range dados.([]interface{}) {

		obj, _ := domain.NewTypeStore(
			item.(map[string]interface{})["description"].(string),
			item.(map[string]interface{})["store_type"].(string),
			craw.List,
			item.(map[string]interface{})["suboptions"].([]interface{}),
		)
		if segment != "" {
			if item.(map[string]interface{})["store_type"].(string) == segment {
				typestore = append(typestore, obj)
			}
		} else {
			typestore = append(typestore, obj)
		}
	}

	craw.TypeStoreRepository.InsertMany(typestore)

	return typestore, nil

}
