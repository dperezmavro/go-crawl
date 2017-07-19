package main

import (
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var regex = regexp.MustCompile("href=\"(https?://)?([a-zA-Z/.%0-9-]+)\"")

func extractLinks(page []byte) ([]string, error) {
	results := regex.FindAllStringSubmatch(string(page), -1)
	var filteredResults []string
	for _, result := range results {
		filteredResults = append(filteredResults, result[2])
	}

	return filteredResults, nil
}

func pollToCrawlChan() {
	log.Println("[+] pollToCrawlChan")
	for {
		var url string
		select {
		case url = <-toCrawl:
			wg.Add(1)

			u, err := formatUrl(url)
			checkErr(err)
			go getPage(u)
		case <-time.After(time.Second * 10):
			wg.Done()
			return
		}
	}
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

func getPage(u *url.URL) {
	defer wg.Done()

	body, err := get(u)
	checkErr(err)
	urls, err := extractLinks(body)
	checkErr(err)

	doneCrawaling <- crawlResult{
		u.String(),
		urls,
	}
}
