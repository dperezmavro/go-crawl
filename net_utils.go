package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return body, nil
}
