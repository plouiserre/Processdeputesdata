package main

import (
	"database/sql"
	"time"
)

type DeputyRepository struct {
	Sql  SqlManager
	Log  LogManager
	Data DataManager
}

func (repository *DeputyRepository) RecordAllDeputyData() {
	repository.RecordDeputyData()
}

//TODO à optimiser plus tard
func (repository *DeputyRepository) RecordDeputyData() {
	mandates := repository.Data.CongressManModel.Mandates

	for _, mandate := range mandates {
		deputy := mandate.Deputy
		if deputy != (Deputy{}) {
			repository.InsertRepository(deputy)
		}
	}
}

//TODO normalement il faudrait que je mette un lien entre mandate et deputy mais pas pour le moment regarder notes du 11/09/21
func (repository *DeputyRepository) InsertRepository(deputy Deputy) {
	if len(deputy.EndDate) > 0 {
		repository.InsertRepositoryWithEndDate(deputy)
	} else {
		repository.InsertRepositoryWithNoEndDate(deputy)
	}
}

func (repository *DeputyRepository) InsertRepositoryWithEndDate(deputy Deputy) {
	var startDate time.Time
	var errConvertStart error
	var endDate time.Time
	var errConvertEnd error
	queryDeputy := "INSERT INTO Deputy(StartDate, EndDate, RefDeputy) VALUES (?,?,?)"

	startDate, errConvertStart = time.Parse(time.RFC3339, deputy.StartDate)
	endDate, errConvertEnd = time.Parse(time.RFC3339, deputy.EndDate)

	if errConvertStart != nil {
		repository.Log.WriteErrorLog("Erreur de convertion  de startDate " + deputy.StartDate + " du deputy " + deputy.RefDeputy)
	} else if errConvertEnd != nil {
		repository.Log.WriteErrorLog("Erreur de convertion  de endDate " + deputy.EndDate + " du deputy " + deputy.RefDeputy)
	} else {
		stmt, db, err := repository.PrepareQuery(queryDeputy)

		if err == nil {
			_, errExec := stmt.Exec(startDate, endDate, deputy.RefDeputy)
			if errExec != nil {
				repository.Log.WriteErrorLog("Erreur exécution requête " + errExec.Error())
			}
		}

		defer db.Close()
	}
}

func (repository *DeputyRepository) InsertRepositoryWithNoEndDate(deputy Deputy) {
	var startDate time.Time
	var errConvertStart error
	queryDeputy := "INSERT INTO Deputy(StartDate, RefDeputy) VALUES (?,?)"

	startDate, errConvertStart = time.Parse(time.RFC3339, deputy.StartDate)

	if errConvertStart != nil {
		repository.Log.WriteErrorLog("Erreur de convertion  de startDate " + deputy.StartDate + " du deputy " + deputy.RefDeputy)
	} else {
		stmt, db, err := repository.PrepareQuery(queryDeputy)

		if err == nil {
			_, errExec := stmt.Exec(startDate, deputy.RefDeputy)
			if errExec != nil {
				repository.Log.WriteErrorLog("Erreur exécution requête " + errExec.Error())
			}
		}

		defer db.Close()
	}
}

func (repository *DeputyRepository) PrepareQuery(query string) (*sql.Stmt, *sql.DB, error) {
	db := repository.Sql.InitDB()

	stmt, err := db.Prepare(query)
	if err != nil {
		repository.Log.WriteErrorLog("Erreur préparation requête " + err.Error())
	}

	return stmt, db, err
}
