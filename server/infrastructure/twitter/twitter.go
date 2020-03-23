package twitter

import (
	"context"
	"github.com/dghubble/go-twitter/twitter"
	"server/domain/service"
	secrets_twitter "server/etc/secrets/twitter"
)

type TwitterService struct {
	Client twitter.Client
}

func NewTwitterService() service.TwitterService {
	config := secrets_twitter.OAuthConfig
	httpClient := config.Client(context.TODO())
	client := twitter.NewClient(httpClient)
	twitterService := TwitterService{*client}
	return &twitterService
}

func (s *TwitterService) Search(option twitter.SearchTweetParams) ([]twitter.Tweet, error) {
	result, _, err := s.Client.Search.Tweets(&option)
	if err != nil {
		return nil, err
	}
	return result.Statuses, nil
}
