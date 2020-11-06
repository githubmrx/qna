package db

import (
	"database/sql"
	"qna/common"
	"qna/model"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var Dbmap = initDb()

func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "db_user:db_pass@tcp(192.168.1.23:3306)/questions_n_answers")
	common.CheckErr(err, "sql.Open failed")
	// err = db.Ping()
	// checkErr(err, "db.Ping() failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(model.User{}, "user").SetKeys(true, "Id")
	dbmap.AddTableWithName(model.Question{}, "question").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	common.CheckErr(err, "Create tables failed")

	return dbmap
}
