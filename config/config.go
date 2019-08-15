package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//GetDB Function
// func GetDB() (db *sql.DB, err error) {
// 	dbDriver := "mysql"
// 	dbuser := "root"
// 	dbpass := ""
// 	dbName := "feedbackdb"

// 	db, err = sql.Open(dbDriver, dbuser+":"+dbpass+"@/"+dbName)
// 	return
// }

//GetCloudDB Function
func GetCloudDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbuser := "root"
	dbpass := ""
	dbName := "heineken"

	db, err = sql.Open(dbDriver, dbuser+":"+dbpass+"@/"+dbName)
	return
}
