package main

import (
	"flag"
	"fmt"
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

func pollToCrawlChan() {
	log.Println("[+] pollToCrawlChan")
	for {
		var ok bool
		var url string
		url, ok = <-toCrawl

		fmt.Printf("OK: %v", ok)
		wg.Add(1)
		go getPage(url)
	}

	wg.Done()
}

func getPage(u string) {
	defer wg.Done()

	log.Printf("[+] URL: %s", u)

	var tempUrl url.URL
	if !tempUrl.IsAbs() {
		tmp, err := url.Parse("http://" + tempUrl.String())
		checkErr(err)

		tempUrl = *tmp
	}

	body, err := get(tempUrl)
	checkErr(err)

	urls, err := extractLinks(body)
	checkErr(err)

	urls = filterExternalLinks(
		urls,
		stripProtocol(crawlUrl),
	)
	urls = stripPrefix(
		urls,
		urlNoProtocol,
	)

	doneCrawaling <- crawlResult{
		tempUrl.String(),
		urls,
	}
}

func processResults() {
	log.Println("[+] processResults")
	for {
		var result = <-doneCrawaling
		log.Printf("[+] Crawling url %s", result.url)
		// for _, u := range result.links {
		// 	tempUrl, err := url.Parse(u)
		// 	checkErr(err)
		// 	if tempUrl.IsAbs() {
		// 		//toCrawl <- *tempUrl
		// 	}
		// }
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
