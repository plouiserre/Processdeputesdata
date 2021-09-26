package main

import "database/sql"

type MandateRepository struct {
	Sql                SqlManager
	Log                LogManager
	Data               DataManager
	DeputyRepository   DeputyRepository
	ElectionRepository ElectionRepository
	MandateId          int64
}

func (mandateRepository *MandateRepository) RecordAllMandates(congressManId int64) {
	model := mandateRepository.Data.CongressManModel
	for _, mandateModel := range model.Mandates {
		if len(mandateModel.EndDate) > 0 {
			mandateRepository.RecordMandateWithEndDate(mandateModel, congressManId)
		} else {
			mandateRepository.RecordMandateWithNoEndDate(mandateModel, congressManId)
		}
		mandateRepository.DeputyRepository.RecordDeputyDatas(mandateRepository.MandateId, mandateModel.Deputy)
		mandateRepository.ElectionRepository.RecordElection(mandateRepository.MandateId, mandateModel.Election)
	}
}

//TODO tout passer en paramètre de la structure après avoir fait la boucle
func (mandateRepository *MandateRepository) RecordMandateWithEndDate(mandateModel Mandate, congressManId int64) {
	queryCongressMan := "INSERT INTO PROCESSDEPUTES.Mandate(MandateUid, TermOffice, TypeOrgane, StartDate, EndDate, Precedence, PrincipleNoming, QualityCode, QualityLabel, QualityLabelSex, RefBody, CongressManId) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"

	stmt, db, err := mandateRepository.PrepareQuery(queryCongressMan)

	if err == nil {
		res, errExec := stmt.Exec(mandateModel.MandateUid, mandateModel.TermOffice, mandateModel.TypeOrgane, mandateModel.StartDate, mandateModel.EndDate, mandateModel.Precedence, mandateModel.PrincipleNomin, mandateModel.QualityCode, mandateModel.QualityLabel, mandateModel.QualityLabelSex, mandateModel.RefBody, congressManId)
		if errExec != nil {
			mandateRepository.Log.WriteErrorLog("Mandate Repository : Erreur exécution requête " + errExec.Error())
		}
		//TODO factoriser cette partie aussi
		mandateId, errGetLastId := res.LastInsertId()
		if errGetLastId != nil {
			mandateRepository.Log.WriteErrorLog("Mandate Repository : Erreur récupérage Id" + errGetLastId.Error())
		}
		mandateRepository.MandateId = mandateId
	}

	defer db.Close()

}

func (mandateRepository *MandateRepository) RecordMandateWithNoEndDate(mandateModel Mandate, congressManId int64) {

	queryCongressMan := "INSERT INTO PROCESSDEPUTES.Mandate(MandateUid, TermOffice, TypeOrgane, StartDate, Precedence, PrincipleNoming, QualityCode, QualityLabel, QualityLabelSex, RefBody, CongressManId) VALUES (?,?,?,?,?,?,?,?,?,?,?)"

	stmt, db, err := mandateRepository.PrepareQuery(queryCongressMan)

	if err == nil {
		res, errExec := stmt.Exec(mandateModel.MandateUid, mandateModel.TermOffice, mandateModel.TypeOrgane, mandateModel.StartDate, mandateModel.Precedence, mandateModel.PrincipleNomin, mandateModel.QualityCode, mandateModel.QualityLabel, mandateModel.QualityLabelSex, mandateModel.RefBody, congressManId)
		if errExec != nil {
			mandateRepository.Log.WriteErrorLog("Mandate Repository : Erreur exécution requête " + errExec.Error())
		}
		//TODO factoriser cette partie aussi
		mandateId, errGetLastId := res.LastInsertId()
		if errGetLastId != nil {
			mandateRepository.Log.WriteErrorLog("Mandate Repository : Erreur récupérage Id" + errGetLastId.Error())
		}
		mandateRepository.MandateId = mandateId
	}

	defer db.Close()
}

//TODO la factoriser avec la deputyrepository
func (mandateRepository *MandateRepository) PrepareQuery(query string) (*sql.Stmt, *sql.DB, error) {
	db := mandateRepository.Sql.InitDB()

	stmt, err := db.Prepare(query)
	if err != nil {
		mandateRepository.Log.WriteErrorLog("Mandate Repository : Erreur préparation requête " + err.Error())
	}

	return stmt, db, err
}
