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
		case <-time.After(timeout):
			wg.Done()
			log.Println("[+] pollToCrawlChan timeout!")
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

		case <-time.After(timeout):
			wg.Done()
			close(toCrawl)
			log.Println("[+] processResults timeout")
			return
		}

	}
}
