package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"strings"
)

func initialise() {
	flag.StringVar(&startingCrawlUrl, "url", "", "A Url to crawl")
	flag.Parse()
}

func main() {
	initialise()

	if startingCrawlUrl == "" {
		flag.Usage()
		os.Exit(1)
	}

	url, err := url.Parse(startingCrawlUrl)
	checkErr(err)

	if !url.IsAbs() {
		startingCrawlUrl = strings.Join(
			[]string{"http://", startingCrawlUrl},
			"",
		)
		url, err = url.Parse(startingCrawlUrl)
		checkErr(err)
	}
	hostName = url.Hostname()
	toCrawl <- startingCrawlUrl

	wg.Add(2)
	go pollToCrawlChan()
	go processResults()

	wg.Wait()

	for v, _ := range urls {
		log.Printf("[+] url : %s", v)
	}
}
