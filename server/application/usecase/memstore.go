package usecase

import (
	"encoding/json"
	"github.com/dghubble/go-twitter/twitter"
	domain "server/domain/model"
	"server/domain/service"
)

type MemStoreUsecase interface {
	// ogp
	AddCacheOgp(ogp *domain.Ogp) error
	GetCacheOgp(url string) (*domain.Ogp, error)
	// tweet
	AddCacheTweet(tweet twitter.Tweet) error
	HasCacheTweet(id string) (bool, error)
	// rate limit
	AddRateLimitRecordOgp(fqdn string, expireSec int) error
	IsRateLimitedRecordOgp(fqdn string) (bool, error)
}

type memStoreUsecase struct {
	Service service.MemStoreService
}

func NewMemStoreUsecase(memStoreService service.MemStoreService) MemStoreUsecase {
	memStoreUsecase := memStoreUsecase{memStoreService}
	return &memStoreUsecase
}

func (u *memStoreUsecase) AddCacheOgp(ogp *domain.Ogp) error {
	field := "ogp"
	key := field + ":" + ogp.URL
	expireSec := 60 * 10 // 10 min
	return u.Service.Add(field, key, ogp, expireSec)
}

func (u *memStoreUsecase) GetCacheOgp(url string) (*domain.Ogp, error) {
	field := "ogp"
	b, err := u.Service.Get(field, url)
	if err != nil || b == nil {
		return nil, err
	}
	ogpJson := b.(string)
	var ogp *domain.Ogp
	err = json.Unmarshal([]byte(ogpJson), &ogp)
	if err != nil {
		return nil, err
	}
	return ogp, nil
}

func (u *memStoreUsecase) AddCacheTweet(tweet twitter.Tweet) error {
	field := "twitter"
	key := field + ":" + tweet.IDStr
	expireSec := 60
	return u.Service.Add(field, key, tweet, expireSec)
}

func (u *memStoreUsecase) HasCacheTweet(id string) (bool, error) {
	field := "twitter"
	return u.Service.HasCache(field, id)
}

func (u *memStoreUsecase) AddRateLimitRecordOgp(fqdn string, expireSec int) error {
	field := "fqdn"
	key := field + ":" + fqdn
	return u.Service.Add(field, key, true, expireSec)
}

func (u *memStoreUsecase) IsRateLimitedRecordOgp(fqdn string) (bool, error) {
	field := "fqdn"
	return u.Service.HasCache(field, fqdn)
}
