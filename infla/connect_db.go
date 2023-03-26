package infla

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yoshikawa-river/ChatApp/config"
)

func NewConnDB(dbInfo *config.DBInfo) (*sql.DB, error) {
	connect := connDBInfo(dbInfo)
	db, err := sql.Open(dbInfo.CONN, connect)

	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db, err
}

func connDBInfo(dbInfo *config.DBInfo) string {
	return fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=true&loc=Local", dbInfo.USER, dbInfo.PASSWORD, dbInfo.DBPROTOCOL, dbInfo.DBNAME)
}
