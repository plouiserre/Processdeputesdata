package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type SqlManager struct {
	Log LogManager
}

//TODO tout mettre en configuration
func (sqlManager *SqlManager) InitDB() (db *sql.DB) {
	db, err := sql.Open("mysql", "ProcessDeputesData:ASimpleP@ssW0rd@/PROCESSDEPUTES")
	if err != nil {
		sqlManager.Log.WriteErrorLog("Erreur connexion " + err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db
}
