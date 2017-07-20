package main

import (
	"flag"
	"log"
	"os"
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

	url, err := formatUrl(startingCrawlUrl)
	checkErr(err)

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
