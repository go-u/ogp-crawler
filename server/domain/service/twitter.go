package service

import (
	"github.com/dghubble/go-twitter/twitter"
)

type TwitterService interface {
	Search(option twitter.SearchTweetParams) ([]twitter.Tweet, error)
}
