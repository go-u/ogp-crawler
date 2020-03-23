package usecase

import (
	domain "server/domain/model"
	"server/domain/repository"
)

type OgpUsecase interface {
	RecordSample(ogp *domain.Ogp) error
}

type ogpUsecase struct {
	Repo repository.OgpRepository
}

func NewOgpUsecase(ogpRepo repository.OgpRepository) OgpUsecase {
	ogpUsecase := ogpUsecase{ogpRepo}
	return &ogpUsecase
}

func (u *ogpUsecase) RecordSample(ogp *domain.Ogp) error {
	err := u.Repo.Record(ogp)
	return err
}
