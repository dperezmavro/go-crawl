package main

import (
	"flag"
	"log"
	"os"
)

var Url string

func initialise() {
	flag.StringVar(&Url, "url", "", "A Url to crawl")
	flag.Parse()
}

func main() {
	initialise()

	if Url == "" {
		flag.Usage()
		os.Exit(1)
	}

	urlNoProtocol := stripProtocol(Url)

	page, err := get(Url)
	checkErr(err)

	urls, err := extractLinks(page)
	checkErr(err)

	urls = filterExternalLinks(
		urls,
		stripProtocol(Url),
	)
	urls = stripPrefix(
		urls,
		urlNoProtocol,
	)

	for _, url := range urls {
		log.Print(url)
	}
}
