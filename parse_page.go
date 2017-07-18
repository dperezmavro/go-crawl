package main

import (
	"log"
	"net/url"
	"regexp"
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
		var ok bool
		var url string
		url, ok = <-toCrawl
		if !ok {
			log.Println("Exiting!")
			break
		}

		wg.Add(1)
		go getPage(url)
	}

	wg.Done()
}

func formatUrl(u string) (*url.URL, error) {
	tempUrl, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	if !tempUrl.IsAbs() {
		tempUrl, err = url.Parse("http://" + u)
		if err != nil {
			return nil, err
		}
	}

	return tempUrl, nil
}

func getPage(u string) {
	defer wg.Done()

	log.Printf("[+] URL: %s", u)

	tempUrl, err := formatUrl(u)
	checkErr(err)

	body, err := get(tempUrl)
	checkErr(err)
	urls, err := extractLinks(body)
	checkErr(err)

	// urls = filterExternalLinks(
	// 	urls,
	// 	stripProtocol(crawlUrl),
	// )

	// urls = stripPrefix(
	// 	urls,
	// 	urlNoProtocol,
	// )

	doneCrawaling <- crawlResult{
		tempUrl.String(),
		urls,
	}
}
