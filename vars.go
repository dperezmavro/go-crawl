package main

import (
	"sync"
	"time"
)

const timeout = time.Second * 5

var startingCrawlUrl string
var hostName string
var toCrawl = make(chan string, 5000)
var doneCrawling = make(chan crawlResult)
var wg sync.WaitGroup
var urls = make(map[string]string)

type crawlResult struct {
	url   string
	links []string
}
