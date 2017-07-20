package main

import (
	"log"
	"time"
)

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
		case <-time.After(time.Second * 5):
			wg.Done()
			return
		}
	}
}

func processResults() {
	log.Println("[+] processResults")
	var result crawlResult
	for {
		select {
		case result = <-doneCrawling:
			for _, u := range result.links {
				tempUrl, err := formatUrl(u)
				checkErr(err)
				storeUrl(tempUrl)
			}

		case <-time.After(time.Second * 5):
			wg.Done()
			close(toCrawl)
			return
		}

	}
}
