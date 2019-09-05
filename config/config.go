package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//GetCloudDB Function
func GetCloudDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	// dbuser := "rtuser"
	// dbpass := "rtuser123"
	dbuser := "root"
	dbpass := ""
	dbName := "heineken"
	db, err = sql.Open(dbDriver, dbuser+":"+dbpass+"@/"+dbName)
	return
}
