package main

type CongressManRepository struct {
	RepositoryManager *RepositoryManager
	MandateRepository MandateRepository
}

func (repository *CongressManRepository) RecordAllCongressManData() {
	congressManId := repository.RecordCongressManData()
	repository.MandateRepository.CongressManId = congressManId
	repository.MandateRepository.RecordAllMandates()
}

func (congressManRepository *CongressManRepository) RecordCongressManData() int64 {
	var lid int64
	repository := congressManRepository.RepositoryManager
	congressMan := repository.Data.CongressManModel
	queryCongressMan := "INSERT INTO PROCESSDEPUTES.Congressman(CongressManUid, Civility, FirstName, LastName, Alpha, Trigramme, BirthDate, BirthCity, BirthDepartment, BirthCountry, JobTitle, CatSocPro, FamSocPro) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"

	db := repository.Sql.InitDB()

	nameRepository := "CongressMan Repository"

	stmt, isOkPrepareRequest := repository.Sql.PrepareRequest(db, queryCongressMan, nameRepository)

	if isOkPrepareRequest {
		res, errExec := stmt.Exec(congressMan.CongressManUid, congressMan.Civility, congressMan.FirstName, congressMan.LastName, congressMan.Alpha, congressMan.Trigramme, congressMan.BirthDate, congressMan.BirthCity, congressMan.BirthDepartment, congressMan.BirthCountry, congressMan.JobTitle, congressMan.CatSocPro, congressMan.FamSocPro)
		if errExec != nil {
			repository.Log.WriteErrorLog("Congressman Repository : Erreur exécution requête " + errExec.Error())
		}

		lid = repository.Sql.GetLastIdInsert(res, nameRepository)
	}

	defer db.Close()

	return lid
}
