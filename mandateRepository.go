package main

import "database/sql"

//TODO d'abord faire electionRepository avant
type MandateRepository struct {
	Sql  SqlManager
	Log  LogManager
	Data DataManager
}

func (mandateRepository *MandateRepository) RecordAllMandates(congressManId int64) {
	//test := strconv.FormatInt(congressManId, 10)
	//mandateRepository.Log.WriteErrorLog("CongressManId " + test)
	model := mandateRepository.Data.CongressManModel
	for _, mandateModel := range model.Mandates {
		if len(mandateModel.EndDate) > 0 {
			mandateRepository.RecordMandateWithEndDate(mandateModel, congressManId)
		} else {
			mandateRepository.RecordMandateWithNoEndDate(mandateModel, congressManId)
		}
	}
}

//TODO factoriser
//TODO tout passer en paramètre de la structure après avoir fait la boucle
func (mandateRepository *MandateRepository) RecordMandateWithEndDate(mandateModel Mandate, congressManId int64) {

	queryCongressMan := "INSERT INTO PROCESSDEPUTES.Mandate(MandateUid, TermOffice, TypeOrgane, StartDate, EndDate, Precedence, PrincipleNoming, QualityCode, QualityLabel, QualityLabelSex, RefBody, CongressManId) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"

	stmt, db, err := mandateRepository.PrepareQuery(queryCongressMan)

	if err == nil {
		_, errExec := stmt.Exec(mandateModel.MandateUid, mandateModel.TermOffice, mandateModel.TypeOrgane, mandateModel.StartDate, mandateModel.EndDate, mandateModel.Precedence, mandateModel.PrincipleNomin, mandateModel.QualityCode, mandateModel.QualityLabel, mandateModel.QualityLabelSex, mandateModel.RefBody, congressManId)
		if errExec != nil {
			mandateRepository.Log.WriteErrorLog("Erreur exécution requête " + errExec.Error())
		}
	}

	defer db.Close()
}

func (mandateRepository *MandateRepository) RecordMandateWithNoEndDate(mandateModel Mandate, congressManId int64) {

	queryCongressMan := "INSERT INTO PROCESSDEPUTES.Mandate(MandateUid, TermOffice, TypeOrgane, StartDate, Precedence, PrincipleNoming, QualityCode, QualityLabel, QualityLabelSex, RefBody, CongressManId) VALUES (?,?,?,?,?,?,?,?,?,?,?)"

	stmt, db, err := mandateRepository.PrepareQuery(queryCongressMan)

	if err == nil {
		_, errExec := stmt.Exec(mandateModel.MandateUid, mandateModel.TermOffice, mandateModel.TypeOrgane, mandateModel.StartDate, mandateModel.Precedence, mandateModel.PrincipleNomin, mandateModel.QualityCode, mandateModel.QualityLabel, mandateModel.QualityLabelSex, mandateModel.RefBody, congressManId)
		if errExec != nil {
			mandateRepository.Log.WriteErrorLog("Erreur exécution requête " + errExec.Error())
		}
	}

	defer db.Close()
}

//TODO la factoriser avec la deputyrepository
func (mandateRepository *MandateRepository) PrepareQuery(query string) (*sql.Stmt, *sql.DB, error) {
	db := mandateRepository.Sql.InitDB()

	stmt, err := db.Prepare(query)
	if err != nil {
		mandateRepository.Log.WriteErrorLog("Erreur préparation requête " + err.Error())
	}

	return stmt, db, err
}
