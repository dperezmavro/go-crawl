package main

import (
	"net/url"
)

func isExternal(u *url.URL) bool {
	return u.Hostname() != hostName
}

func filterExternalLinks(urls []string, domain string) []string {
	var result []string
	for _, u := range urls {
		tempUrl, err := url.Parse(u)
		checkErr(err)
		if !isExternal(tempUrl) {
			result = append(result, tempUrl.String())
		}
	}
	return result
}
