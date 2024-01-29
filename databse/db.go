package database

import (
	"database/sql"
	"ginsdgo/utils"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func InitDB() {
	var err error 
	db, err := sql.Open("mysql", utils.DB_URL)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
}

func init()	{
	InitDB()
}

func GetDB() *sql.DB {
	return db
}