package main

import (
	"net/url"
	"strings"
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

func formatUrl(u string) (*url.URL, error) {
	tempUrl, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(tempUrl.String(), "/") {
		tempUrl.Host = hostName
		tempUrl.Scheme = "http"
	}

	if !tempUrl.IsAbs() {
		tempUrl, err = url.Parse("http://" + u)
		if err != nil {
			return nil, err
		}
	}

	return tempUrl, nil
}
