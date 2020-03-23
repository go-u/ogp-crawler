package main

import (
	"server/application/usecase"
	"server/crawler"
	"server/infrastructure/memstore"
	"server/infrastructure/store"
	"server/infrastructure/twitter"
)

func getCrawlerConfig(option Options) crawler.Config {
	projectID := getProjectId()

	// infra
	sqlHandler := store.NewSqlHandler(projectID)

	// repository & service
	ogpRepository := store.NewOgpRepository(*sqlHandler)
	statRepository := store.NewStatRepository(*sqlHandler)
	memStoreService := memstore.NewMemStoreService(6379)
	twitterService := twitter.NewTwitterService()

	// usecase
	ogpUsecase := usecase.NewOgpUsecase(ogpRepository)
	statUsecase := usecase.NewStatUsecase(statRepository)
	memStoreUsecase := usecase.NewMemStoreUsecase(memStoreService)
	twitterUsecase := usecase.NewTwitterUsecase(twitterService)

	config := crawler.Config{
		HostName:        getHostName(),
		OgpUsecase:      ogpUsecase,
		StatUsecase:     statUsecase,
		MemStoreUsecase: memStoreUsecase,
		TwitterUsecase:  twitterUsecase,

		// option
		RateLimitDurationOgp: option.RateLimit.Ogp,
	}
	return config
}
