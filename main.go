package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"
)

var crawlUrl string
var hostName string
var urlNoProtocol string
var toCrawl = make(chan string, 5000)
var doneCrawaling = make(chan crawlResult)
var wg sync.WaitGroup

type crawlResult struct {
	url   string
	links []string
}

func initialise() {
	flag.StringVar(&crawlUrl, "url", "", "A Url to crawl")
	flag.Parse()
}

func processResults() {
	log.Println("[+] processResults")
	for {
		var result = <-doneCrawaling
		log.Printf("[+] Done crawling url %s, links : %v", result.url, result.links)
		for _, u := range result.links {
			tempUrl, err := url.Parse(u)
			checkErr(err)
			if tempUrl.IsAbs() {
				//toCrawl <- *tempUrl
			}
		}
	}

	wg.Done()
}

func main() {
	initialise()

	if crawlUrl == "" {
		flag.Usage()
		os.Exit(1)
	}

	url, err := url.Parse(crawlUrl)
	checkErr(err)

	hostName = url.Hostname()
	if !url.IsAbs() {
		crawlUrl = strings.Join([]string{"http://", crawlUrl}, "")
	}

	urlNoProtocol = stripProtocol(crawlUrl)

	toCrawl <- url.String()

	wg.Add(2)
	go pollToCrawlChan()
	go processResults()

	wg.Wait()

	// for _, url := range urls {
	// 	log.Print(url)
	// }
}
