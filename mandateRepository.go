package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type MandateRepository struct {
	RepositoryManager  *RepositoryManager
	DeputyRepository   DeputyRepository
	ElectionRepository ElectionRepository
	MandateId          int64
	CongressManId      int64
}

func (mandateRepository *MandateRepository) RecordAllMandates() {
	repository := mandateRepository.RepositoryManager
	model := repository.Data.CongressManModel
	for _, mandateModel := range model.Mandates {
		if len(mandateModel.EndDate) > 0 {
			mandateRepository.RecordMandateWithEndDate(mandateModel)
		} else {
			mandateRepository.RecordMandateWithNoEndDate(mandateModel)
		}
		mandateRepository.DeputyRepository.RecordDeputyDatas(mandateRepository.MandateId, mandateModel.Deputy)
		mandateRepository.ElectionRepository.RecordElection(mandateRepository.MandateId, mandateModel.Election)
	}
}

func (mandateRepository *MandateRepository) RecordMandateWithEndDate(mandateModel MandateModel) {
	queryMandate := "INSERT INTO PROCESSDEPUTES.Mandate(MandateUid, TermOffice, TypeOrgane, StartDate, EndDate, Precedence, PrincipleNoming, QualityCode, QualityLabel, QualityLabelSex, RefBody, CongressManId) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"
	repository := mandateRepository.RepositoryManager
	db := repository.Sql.InitDB()
	nameRepository := "Mandate Repository"

	stmt, isOk := repository.Sql.PrepareRequest(db, queryMandate, nameRepository)

	if isOk {
		res, errExec := stmt.Exec(mandateModel.MandateUid, mandateModel.TermOffice, mandateModel.TypeOrgane, mandateModel.StartDate, mandateModel.EndDate, mandateModel.Precedence, mandateModel.PrincipleNomin, mandateModel.QualityCode, mandateModel.QualityLabel, mandateModel.QualityLabelSex, mandateModel.RefBody, mandateRepository.CongressManId)
		if errExec != nil {
			repository.Log.WriteErrorLog("Mandate Repository : Erreur exécution requête " + errExec.Error())
		}
		mandateId := repository.Sql.GetLastIdInsert(res, nameRepository)

		mandateRepository.MandateId = mandateId
	}

	defer db.Close()

}

func (mandateRepository *MandateRepository) RecordMandateWithNoEndDate(mandateModel MandateModel) {
	queryMandate := "INSERT INTO PROCESSDEPUTES.Mandate(MandateUid, TermOffice, TypeOrgane, StartDate, Precedence, PrincipleNoming, QualityCode, QualityLabel, QualityLabelSex, RefBody, CongressManId) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
	repository := mandateRepository.RepositoryManager
	db := repository.Sql.InitDB()
	nameRepository := "Mandate Repository"

	stmt, isOk := repository.Sql.PrepareRequest(db, queryMandate, nameRepository)

	if isOk {
		res, errExec := stmt.Exec(mandateModel.MandateUid, mandateModel.TermOffice, mandateModel.TypeOrgane, mandateModel.StartDate, mandateModel.Precedence, mandateModel.PrincipleNomin, mandateModel.QualityCode, mandateModel.QualityLabel, mandateModel.QualityLabelSex, mandateModel.RefBody, mandateRepository.CongressManId)
		if errExec != nil {
			repository.Log.WriteErrorLog("Mandate Repository : Erreur exécution requête " + errExec.Error())
		}
		mandateId := repository.Sql.GetLastIdInsert(res, nameRepository)

		mandateRepository.MandateId = mandateId
	}

	defer db.Close()
}
