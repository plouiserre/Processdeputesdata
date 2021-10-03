package main

import (
	"time"
)

type DeputyRepository struct {
	RepositoryManager *RepositoryManager
	MandateId         int64
	utility           Utility
}

func (repository *DeputyRepository) RecordDeputyDatas(mandateId int64, deputy DeputyModel) {
	repository.MandateId = mandateId

	if deputy != (DeputyModel{}) {
		repository.InsertDeputy(deputy)
	}
}

func (repository *DeputyRepository) InsertDeputy(deputy DeputyModel) {
	if len(deputy.EndDate) > 0 {
		repository.InsertDeputyWithEndDate(deputy)
	} else {
		repository.InsertDeputyWithNoEndDate(deputy)
	}
}

func (deputyRepository *DeputyRepository) InsertDeputyWithEndDate(deputy DeputyModel) {
	var startDate time.Time
	var endDate time.Time
	repository := deputyRepository.RepositoryManager
	db := repository.Sql.InitDB()
	nameRepository := "Deputy Repository : "
	queryDeputy := "INSERT INTO Deputy(StartDate, EndDate, RefDeputy, MandateId) VALUES (?,?,?,?)"
	errorMessageConvertStart := "Deputy Repository : Erreur de convertion  de startDate " + deputy.StartDate + " du deputy " + deputy.RefDeputy
	errorMessageConvertEnd := "Deputy Repository : Erreur de convertion  de endDate " + deputy.EndDate + " du deputy " + deputy.RefDeputy

	startDate, resultConvertStart := deputyRepository.utility.ConvertStringToTime(deputy.StartDate, errorMessageConvertStart)
	endDate, resultConvertEnd := deputyRepository.utility.ConvertStringToTime(deputy.EndDate, errorMessageConvertEnd)

	if resultConvertStart && resultConvertEnd {
		stmt, isOk := repository.Sql.PrepareRequest(db, queryDeputy, nameRepository)

		if isOk {
			_, errExec := stmt.Exec(startDate, endDate, deputy.RefDeputy, deputyRepository.MandateId)
			if errExec != nil {
				repository.Log.WriteErrorLog("Deputy Repository : Erreur exécution requête " + errExec.Error())
			}
		}

		defer db.Close()
	}
}

func (deputyRepository *DeputyRepository) InsertDeputyWithNoEndDate(deputy DeputyModel) {
	var startDate time.Time
	repository := deputyRepository.RepositoryManager
	db := repository.Sql.InitDB()
	nameRepository := "Deputy Repository : "
	queryDeputy := "INSERT INTO Deputy(StartDate, RefDeputy, MandateId) VALUES (?,?,?)"
	errorMessageConvertStart := "Deputy Repository : Erreur de convertion  de startDate " + deputy.StartDate + " du deputy " + deputy.RefDeputy

	startDate, resultConvertStart := deputyRepository.utility.ConvertStringToTime(deputy.StartDate, errorMessageConvertStart)

	if resultConvertStart {
		stmt, isOk := repository.Sql.PrepareRequest(db, queryDeputy, nameRepository)

		if isOk {
			_, errExec := stmt.Exec(startDate, deputy.RefDeputy, deputyRepository.MandateId)
			if errExec != nil {
				repository.Log.WriteErrorLog("Deputy Repository : Erreur exécution requête " + errExec.Error())
			}
		}

		defer db.Close()
	}
}
