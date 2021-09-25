package main

/*
type ElectionRepository struct {
	Sql  SqlManager
	Log  LogManager
	Data DataManager
}

func (electionRepository *ElectionRepository) RecordElection() {
	congressMan := electionRepository.Data.CongressManModel
	queryElection := "INSERT INTO PROCESSDEPUTES.Election(MandateCause, Region, TypeRegion, Department, DepartmentNum, DisctrictNum) VALUES (?,?,?,?,?,?)"

	db := electionRepository.Sql.InitDB()

	stmt, err := db.Prepare((queryElection))
	if err != nil {
		electionRepository.Log.WriteErrorLog("Erreur préparation requête " + err.Error())
	}

	_, errExec := stmt.Exec(congressMan.CongressManUid, congressMan.Civility, congressMan.FirstName, congressMan.LastName, congressMan.Alpha, congressMan.Trigramme, congressMan.BirthDate, congressMan.BirthCity, congressMan.BirthDepartment, congressMan.BirthCountry, congressMan.JobTitle, congressMan.CatSocPro, congressMan.FamSocPro)
	if errExec != nil {
		electionRepository.Log.WriteErrorLog("Erreur exécution requête " + errExec.Error())
	}

	defer db.Close()
}*/
