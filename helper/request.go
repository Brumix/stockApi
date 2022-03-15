package helper

import (
	"io/ioutil"
	"net/http"
)

func GETRequest(url string) ([]byte, error) {
	response, errRequest := http.Get(url)
	defer response.Body.Close()
	if errRequest != nil {
		return nil, errRequest
	}
	body, errIoutil := ioutil.ReadAll(response.Body)
	if errIoutil != nil {
		return nil, errIoutil
	}
	return body, nil
}
