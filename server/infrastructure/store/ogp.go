package store

import (
	"context"
	"github.com/volatiletech/sqlboiler/boil"
	domain "server/domain/model"
	"server/domain/repository"
	"server/infrastructure/store/models"
)

func NewOgpRepository(sqlHandler SqlHandler) repository.OgpRepository {
	ogpRepository := OgpStore{sqlHandler}
	return &ogpRepository
}

type OgpStore struct {
	SqlHandler
}

func (s *OgpStore) Record(ogp *domain.Ogp) error {
	o := models.Ogp{
		Date:    ogp.Date,
		FQDN:    ogp.FQDN,
		Host:    ogp.HostName,
		TweetID: ogp.TweetID,
		Type:    ogp.Type,
		Lang:    ogp.Lang,
	}
	err := o.Insert(context.Background(), s.Conn, boil.Blacklist())
	return err
}
