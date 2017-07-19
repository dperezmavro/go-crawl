package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

var startingCrawlUrl string
var hostName string
var toCrawl = make(chan string, 5000)
var doneCrawaling = make(chan crawlResult)
var wg sync.WaitGroup
var urls = make(map[string]string)

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
	var result crawlResult
	for {
		select {
		case result = <-doneCrawaling:
			for _, u := range result.links {
				tempUrl, err := formatUrl(u)
				checkErr(err)

				if !isExternal(tempUrl) {
					if urls[tempUrl.String()] == "" {
						urls[tempUrl.String()] = "done"
						toCrawl <- tempUrl.String()
					} else {
						log.Printf("[*] Ignoring existing url %s", tempUrl.String())
					}
				} else {
					log.Printf("[*] Ignoring external URL: %s", tempUrl.String())
				}
			}

		case <-time.After(time.Second * 5):
			wg.Done()
			return
		}

	}
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
