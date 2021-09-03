package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// SendRequestAndParseResponse ...
func SendRequestAndParseResponse(mtd, url string, payload, respObj interface{}) error {
	resp, err := sendRequest(mtd, url, payload)
	if err != nil {
		log.Println("Error occured while making request", err)
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(data))

	// var d respObj
	json.Unmarshal(data, respObj)

	// err = json.NewDecoder(resp.Body).Decode(respObj)
	if err != nil {
		log.Println("Error occured while parsing response body", err)
	}

	return err
}

func sendRequest(mtd, url string, payload interface{}) (*http.Response, error) {
	var req *http.Request
	var err error

	if payload != nil {
		body, err := json.Marshal(payload)
		if err != nil {
			log.Println("Error marshalling request payload: ", err)
			return nil, err
		}
		req, err = http.NewRequest(mtd, url, bytes.NewBuffer(body))
	} else {
		req, err = http.NewRequest(mtd, url, nil)
	}

	if err != nil {
		log.Println("Error occured while creating request", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	return client.Do(req)
}
