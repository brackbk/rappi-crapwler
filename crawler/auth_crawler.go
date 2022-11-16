package crawler

import (
	"gitlab.com/eiprice/spiders/rappi/utils"
)

var Bearer string

func Auth() (string, error) {
	payload := `{"headers":{"normalizedNames":{},"lazyUpdate":null},"grant_type":"guest"}`
	url := "https://services.rappi.com.br/api/auth/guest_access_token"

	Bearer, err := utils.Request("POST", url, payload, nil)

	if err != nil || Bearer == nil {
		return "", err
	}

	return Bearer.(map[string]interface{})["access_token"].(string), nil

}
