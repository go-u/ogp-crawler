package main

import (
	"github.com/volatiletech/sqlboiler/boil"
	"log"
	"server/crawler"
	"time"
)

func init() {
	boil.DebugMode = false
}

func main() {
	// read config.yaml
	option := getOptions()

	// Dependency Injection
	config := getCrawlerConfig(option)
	c := crawler.NewCrawler(&config)

	i := 0
	for {
		// start crawler
		c.SearchAndCrawl()
		log.Println("count: ", i)
		time.Sleep(time.Second * time.Duration(option.RateLimit.Twitter)) // default 10sec
		i += 1
	}
}
