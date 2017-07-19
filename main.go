package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"
)

var startingCrawlUrl string
var hostName string
var toCrawl = make(chan string, 5000)
var doneCrawaling = make(chan crawlResult)
var wg sync.WaitGroup
var urls map[string]string

type crawlResult struct {
	url   string
	links []string
}

func initialise() {
	flag.StringVar(&startingCrawlUrl, "url", "", "A Url to crawl")
	flag.Parse()
}

func processResults() {
	log.Println("[+] processResults")
	for {
		result, ok := <-doneCrawaling
		if !ok {
			break
		}

		for _, u := range result.links {
			tempUrl, err := formatUrl(u)
			checkErr(err)

			if !isExternal(tempUrl) {
				toCrawl <- tempUrl.String()
			} else {
				log.Printf("[*] Ignoring url %s", tempUrl.String())
			}

		}
	}

	wg.Done()
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
	log.Printf("[+] Hostname is %s %s", hostName, url.Hostname())

	toCrawl <- startingCrawlUrl

	wg.Add(2)
	go pollToCrawlChan()
	go processResults()

	wg.Wait()

	for _, url := range urls {
		log.Print(url)
	}
}
