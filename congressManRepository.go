package main

type congressManRepository struct {
	Sql  SqlManager
	Log  LogManager
	Data DataManager
}

func (repository *congressManRepository) RecordAllCongressManData() {
	repository.RecordCongressManData()
}

func (repository *congressManRepository) RecordCongressManData() {
	congressMan := repository.Data.CongressManModel
	queryCongressMan := "INSERT INTO PROCESSDEPUTES.Congressman(CongressManUid, Civility, FirstName, LastName, Alpha, Trigramme, BirthDate, BirthCity, BirthDepartment, BirthCountry, JobTitle, CatSocPro, FamSocPro) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"

	db := repository.Sql.InitDB()

	stmt, err := db.Prepare((queryCongressMan))
	if err != nil {
		repository.Log.WriteErrorLog("Erreur préparation requête " + err.Error())
	}

	_, errExec := stmt.Exec(congressMan.CongressManUid, congressMan.Civility, congressMan.FirstName, congressMan.LastName, congressMan.Alpha, congressMan.Trigramme, congressMan.BirthDate, congressMan.BirthCity, congressMan.BirthDepartment, congressMan.BirthCountry, congressMan.JobTitle, congressMan.CatSocPro, congressMan.FamSocPro)
	if errExec != nil {
		repository.Log.WriteErrorLog("Erreur exécution requête " + errExec.Error())
	}

	defer db.Close()
}
