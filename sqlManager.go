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

func (sqlManager *SqlManager) GetLastIdInsert(res sql.Result, repositoryName string) int64 {
	lid, errGetLastId := res.LastInsertId()
	if errGetLastId != nil {
		sqlManager.Log.WriteErrorLog(repositoryName + ": Erreur récupérage Id" + errGetLastId.Error())
	}
	return lid
}

//TODO renvoyer aussi un boolean pour continuer ou pas
func (sqlManager *SqlManager) PrepareRequest(db *sql.DB, query string, repositoryName string) (*sql.Stmt, bool) {
	isOk := true
	stmt, err := db.Prepare((query))
	if err != nil {
		sqlManager.Log.WriteErrorLog(repositoryName + ": Erreur préparation requête " + err.Error())
		isOk = false
	}
	return stmt, isOk
}
