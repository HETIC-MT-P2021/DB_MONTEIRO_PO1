package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var USER_DB = "root"
var PASS_DB = "root"
var IP_DB = "db:3306"
var NAME_DB = "classicmodels"

var DB *sql.DB

func ConnectToBDD() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", USER_DB, PASS_DB, IP_DB, NAME_DB)
	DB, _ = sql.Open("mysql", dsn)
}
