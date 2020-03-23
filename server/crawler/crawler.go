package crawler

import (
	"github.com/dghubble/go-twitter/twitter"
	"log"
	"server/application/usecase"
	"strings"
)

type Config struct {
	HostName             string // crawlerを起動しているServer名
	RateLimitDurationOgp int
	OgpUsecase           usecase.OgpUsecase
	StatUsecase          usecase.StatUsecase
	MemStoreUsecase      usecase.MemStoreUsecase
	TwitterUsecase       usecase.TwitterUsecase
}

type Option struct {
	SearchOption twitter.SearchTweetParams
	RateLimitOgp int
}

type Crawler struct {
	*Config
}

func NewCrawler(config *Config) *Crawler {
	crawler := &Crawler{
		Config: config,
	}
	return crawler
}

func (c *Crawler) SearchAndCrawl() {
	p := twitter.SearchTweetParams{
		Query:     "filter:links",
		TweetMode: "extended",
		Count:     100, // max 100
	}

	tweetsWithLink, err := c.TwitterUsecase.Search(p)
	if err != nil {
		log.Println(err)
	}
	for _, tweet := range tweetsWithLink {
		isContainLinks := len(tweet.Entities.Urls) > 0
		isNotInternalLink := isContainLinks && !strings.HasPrefix(tweet.Entities.Urls[0].ExpandedURL, "https://twitter.com")
		if isContainLinks && isNotInternalLink {
			// 最大でおおよそ💯100並列(検索結果が最大100件)
			go c.Crawl(tweet)
		}
	}
}
