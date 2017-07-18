package main

import (
	"net/url"
	"strings"
)

func filterExternalLinks(urls []string, domain string) []string {
	var result []string
	for _, u := range urls {
		tempUrl, err := url.Parse(u)
		checkErr(err)
		if tempUrl.Hostname() == hostName {
			result = append(result, tempUrl.String())
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
		strings.TrimPrefix(url, "http://"),
		"https://",
	)
}
