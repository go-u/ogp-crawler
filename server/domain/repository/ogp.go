package repository

import (
	domain "server/domain/model"
)

type OgpRepository interface {
	Record(ogp *domain.Ogp) error
}
