package store

import (
	"database/sql"
	"fmt"
	"log"
	domain "server/domain/model"
	"server/domain/repository"
	"server/infrastructure/store/models"
)

func NewStatRepository(sqlHandler SqlHandler) repository.StatRepository {
	// プリペアステートメントを登録
	query := fmt.Sprintf("INSERT INTO stat (date, fqdn, host, count, title, description, image, type, lang) " +
		"VALUES (NOW(), ?, ?, 1, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE count = count + 1;")
	prepareStmtStat, err := sqlHandler.Conn.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	statRepository := StatStore{
		sqlHandler,
		prepareStmtStat,
	}
	return &statRepository
}

type StatStore struct {
	SqlHandler
	prepareStmt *sql.Stmt
}

func (s *StatStore) Record(ogp *domain.Ogp) error {
	stat := models.Stat{
		FQDN:        ogp.FQDN,
		Host:        ogp.HostName,
		Title:       ogp.Title,
		Description: ogp.Description,
		Image:       ogp.Image,
		Type:        ogp.Type,
		Lang:        ogp.Lang,
	}

	_, err := s.prepareStmt.Exec(
		stat.FQDN,
		stat.Host,
		stat.Title,
		stat.Description,
		stat.Image,
		stat.Type,
		stat.Lang,
	)

	return err
}
