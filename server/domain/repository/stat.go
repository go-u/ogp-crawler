package repository

import (
	domain "server/domain/model"
)

type StatRepository interface {
	Record(ogp *domain.Ogp) error
}
