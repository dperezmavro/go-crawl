package main

import (
	"regexp"
)

func extractLinks(page []byte) ([]string, error) {
	regexp, err := regexp.Compile("href=\"(https?://)?([a-zA-Z/.%0-9-]+)\"")
	if err != nil {
		return nil, err
	}

	results := regexp.FindAllStringSubmatch(string(page), -1)
	var filteredResults []string
	for _, result := range results {
		filteredResults = append(filteredResults, result[2])
	}

	return filteredResults, nil
}
