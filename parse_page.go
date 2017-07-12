package main

import "regexp"

var regex = regexp.MustCompile("href=\"(https?://)?([a-zA-Z/.%0-9-]+)\"")

func extractLinks(page []byte) ([]string, error) {
	results := regex.FindAllStringSubmatch(string(page), -1)
	var filteredResults []string
	for _, result := range results {
		filteredResults = append(filteredResults, result[2])
	}

	return filteredResults, nil
}
