package main

import (
	"database/sql"
	"time"
)

type DeputyRepository struct {
	Sql       SqlManager
	Log       LogManager
	Data      DataManager
	MandateId int64
}

func (repository *DeputyRepository) RecordDeputyDatas(mandateId int64, deputy Deputy) {
	repository.MandateId = mandateId

	if deputy != (Deputy{}) {
		repository.InsertDeputy(deputy)
	}
}

func (repository *DeputyRepository) InsertDeputy(deputy Deputy) {
	if len(deputy.EndDate) > 0 {
		repository.InsertDeputyWithEndDate(deputy)
	} else {
		repository.InsertDeputyWithNoEndDate(deputy)
	}
}

func (repository *DeputyRepository) InsertDeputyWithEndDate(deputy Deputy) {
	var startDate time.Time
	var errConvertStart error
	var endDate time.Time
	var errConvertEnd error
	queryDeputy := "INSERT INTO Deputy(StartDate, EndDate, RefDeputy, MandateId) VALUES (?,?,?,?)"

	//TODO mettre dans une méthode la partie externalsiation
	startDate, errConvertStart = time.Parse(time.RFC3339, deputy.StartDate)
	endDate, errConvertEnd = time.Parse(time.RFC3339, deputy.EndDate)

	if errConvertStart != nil {
		repository.Log.WriteErrorLog("Deputy Repository : Erreur de convertion  de startDate " + deputy.StartDate + " du deputy " + deputy.RefDeputy)
	} else if errConvertEnd != nil {
		repository.Log.WriteErrorLog("Deputy Repository : Erreur de convertion  de endDate " + deputy.EndDate + " du deputy " + deputy.RefDeputy)
	} else {
		stmt, db, err := repository.PrepareQuery(queryDeputy)

		if err == nil {
			_, errExec := stmt.Exec(startDate, endDate, deputy.RefDeputy, repository.MandateId)
			if errExec != nil {
				repository.Log.WriteErrorLog("Deputy Repository : Erreur exécution requête " + errExec.Error())
			}
		}

		defer db.Close()
	}
}

func (repository *DeputyRepository) InsertDeputyWithNoEndDate(deputy Deputy) {
	var startDate time.Time
	var errConvertStart error
	queryDeputy := "INSERT INTO Deputy(StartDate, RefDeputy, MandateId) VALUES (?,?,?)"

	startDate, errConvertStart = time.Parse(time.RFC3339, deputy.StartDate)

	if errConvertStart != nil {
		repository.Log.WriteErrorLog("Deputy Repository : Erreur de convertion  de startDate " + deputy.StartDate + " du deputy " + deputy.RefDeputy)
	} else {
		stmt, db, err := repository.PrepareQuery(queryDeputy)

		if err == nil {
			_, errExec := stmt.Exec(startDate, deputy.RefDeputy, repository.MandateId)
			if errExec != nil {
				repository.Log.WriteErrorLog("Deputy Repository : Erreur exécution requête " + errExec.Error())
			}
		}

		defer db.Close()
	}
}

func (repository *DeputyRepository) PrepareQuery(query string) (*sql.Stmt, *sql.DB, error) {
	db := repository.Sql.InitDB()

	stmt, err := db.Prepare(query)
	if err != nil {
		repository.Log.WriteErrorLog("Deputy Repository : Erreur préparation requête " + err.Error())
	}

	return stmt, db, err
}
