package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func get(url *url.URL) ([]byte, error) {
	resp, err := http.Get(url.String())
	if err != nil {
		log.Printf("[-] get(1): %q", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[-] get(2): %q", err)
		return nil, err
	}

	return body, nil
}

func getPage(u *url.URL) {
	defer wg.Done()

	body, err := get(u)
	checkErr(err)
	urls, err := extractLinks(body)
	checkErr(err)

	doneCrawling <- crawlResult{
		u.String(),
		urls,
	}
}
