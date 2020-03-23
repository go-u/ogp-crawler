package usecase

import (
	domain "server/domain/model"
	"server/domain/repository"
)

type StatUsecase interface {
	Record(ogp *domain.Ogp) error
}

type statUsecase struct {
	Repo repository.StatRepository
}

func NewStatUsecase(statRepo repository.StatRepository) StatUsecase {
	statUsecase := statUsecase{statRepo}
	return &statUsecase
}

func (u *statUsecase) Record(ogp *domain.Ogp) error {
	err := u.Repo.Record(ogp)
	return err
}
