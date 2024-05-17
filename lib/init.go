package lib

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
	dbUser   = "root"
	dbPass   = "1234"
	dbName   = "gocrud"
)

func makeBaseMigration() {

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
}
