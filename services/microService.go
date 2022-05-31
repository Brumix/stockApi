package services

import (
	"encoding/json"
	"fmt"
	"os"
	"stockApi/helper"
	"stockApi/models"
)

func buildURL(path string) string {

	return fmt.Sprintf("%v", os.Getenv("MICROSERVICE_URL")+":"+os.Getenv("FETCH_PORT")+path)
}

func HelloMicroService(res *models.Hello) error {
	response, err := helper.Hello(buildURL("/"))
	if err != nil {
		return err
	}
	errJson := json.Unmarshal(response, &res)
	if errJson != nil {
		return errJson
	}
	return nil
}

func ParseReponse(path string, method string, request interface{}) (interface{}, error) {
	response, err := helper.Request(buildURL(path), method, request)
	if err != nil {
		return nil, err
	}
	var res interface{}
	errJson := json.Unmarshal(response, &res)
	if errJson != nil {
		return nil, errJson
	}
	return res, nil
}
