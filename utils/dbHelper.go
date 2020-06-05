package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
)

var (
	// DBCon is the connection handle
	// for the database
	DBCon *sql.DB
)

func InitDB(driver string, host string, port int32, name string, user string, password string) *sql.DB {
	con, err := sql.Open(driver, fmt.Sprintf("%s:%s@/%s", user, password, name))

	if err != nil {
		panic(err.Error())
	} else {
		log.Info("DB Connection has been initialized")
	}

	DBCon = con
	return DBCon
}
