package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type SqlManager struct {
	Log LogManager
	//db  *sql.DB
}

//TODO tout mettre en configuration
/*func (sqlManager *SqlManager) Connexion() {
	db, err := sql.Open("mysql", "ProcessDeputesData:ASimpleP@ssW0rd@/PROCESSDEPUTES")
	sqlManager.db = db
	if err != nil {
		sqlManager.Log.WriteErrorLog("Erreur connexion " + err.Error())
	}

	sqlManager.db.SetConnMaxLifetime(time.Minute * 3)
	sqlManager.db.SetMaxIdleConns(10)
	sqlManager.db.SetMaxOpenConns(10)
	defer sqlManager.db.Close()
}*/

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

/*func (sqlManager *SqlManager) InsertData(insertQuery string, args ...interface{}) {
	stmt, err := sqlManager.db.Prepare((insertQuery))
	if err != nil {
		sqlManager.Log.WriteErrorLog("Erreur préparation requête " + err.Error())
	}
	stmt.Exec(args)
	//TODO améliorer cette ligne
	//sqlManager.Log.WriteInfoLog("Request " + logQuery)
	defer sqlManager.db.Close()
}*/
