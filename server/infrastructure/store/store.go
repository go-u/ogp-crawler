package store

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	secrets_cloudsql "server/etc/secrets/cloudsql"
	"strings"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler(projectID string) *SqlHandler {
	source := getSource(projectID)
	con := createConnection(source)
	sqlHandler := &SqlHandler{
		Conn: con,
	}
	return sqlHandler
}

func getSource(projectID string) string {
	Source := ""
	if strings.Contains(projectID, "prd") {
		Source = secrets_cloudsql.DB_Prd
	} else if strings.Contains(projectID, "stg") {
		Source = secrets_cloudsql.DB_Stg
	} else if strings.Contains(projectID, "test") {
		Source = secrets_cloudsql.DB_Test
	} else {
		Source = secrets_cloudsql.DB_Local
	}
	if Source == "" {
		log.Fatal("can't find db source")
	}
	return Source
}

func createConnection(Source string) *sql.DB {
	con, err := sql.Open("mysql", Source)
	// https://blog.nownabe.com/2017/01/16/570.html#accessing-the-database
	// defer con.Close()
	if err != nil {
		log.Fatal("DB Open Error: ", err)
	}
	log.Println("DB initialized")
	return con
}
