package services

import (
	"encoding/json"
	"fmt"
	"os"
	"stockApi/helper"
	"stockApi/models"
)

func buildURL(path string) string {

	return fmt.Sprintf("http://%v", os.Getenv("MICROSERVICE_URL")+path)
}

func HelloMicroService(res *models.Hello) error {
	response, err := helper.GETRequest(buildURL("/"))
	if err != nil {
		return err
	}
	errJson := json.Unmarshal(response, &res)
	if errJson != nil {
		return errJson
	}
	return nil
}
