package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Hello(url string) ([]byte, error) {
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
func Request(url string, method string, body interface{}) ([]byte, error) {
	client := &http.Client{}
	postBody, _ := json.Marshal(body)
	requestBody := bytes.NewBuffer(postBody)
	request, errRequest := http.NewRequest(method, url, requestBody)
	request.Header.Set("Content-type", "application/json")
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if errRequest != nil {
		return nil, errRequest
	}

	bodyResponse, errIoutil := ioutil.ReadAll(response.Body)
	if errIoutil != nil {
		return nil, errIoutil
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf(string(bodyResponse))
	}
	return bodyResponse, nil
}
