package usecase

import (
	"github.com/dghubble/go-twitter/twitter"
	"server/domain/service"
)

type TwitterUsecase interface {
	Search(twitter.SearchTweetParams) ([]twitter.Tweet, error)
}

type twitterUsecase struct {
	Service service.TwitterService
}

func NewTwitterUsecase(twitterService service.TwitterService) TwitterUsecase {
	twitterUsecase := twitterUsecase{twitterService}
	return &twitterUsecase
}

func (u *twitterUsecase) Search(option twitter.SearchTweetParams) ([]twitter.Tweet, error) {
	return u.Service.Search(option)
}
