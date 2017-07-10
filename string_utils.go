package main

import (
	"strings"
)

func filterExternalLinks(urls []string, domain string) []string {
	var result []string
	for _, url := range urls {
		if strings.HasPrefix(url, domain) || strings.HasPrefix(url, "/") {
			result = append(result, url)
		}
	}
	return result
}

func stripPrefix(list []string, prefix string) []string {
	var result []string
	for _, str := range list {
		result = append(result, strings.TrimPrefix(str, prefix))
	}

	return result
}

func stripProtocol(url string) string {
	return strings.TrimPrefix(
		strings.TrimPrefix(Url, "http://"),
		"https://",
	)
}
