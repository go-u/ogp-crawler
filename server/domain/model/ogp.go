package model

import (
	"time"
)

type Ogp struct {
	HostName    string
	Date        time.Time
	FQDN        string
	URL         string
	Title       string
	Description string
	Image       string
	Type        string
	Lang        string
	TweetID     int64
}
