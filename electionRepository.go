package main

type ElectionRepository struct {
	Sql  SqlManager
	Log  LogManager
	Data DataManager
}

func (electionRepository *ElectionRepository) RecordElection(mandateId int64, election Election) {
	if election != (Election{}) {
		queryElection := "INSERT INTO PROCESSDEPUTES.Election(MandateCause, Region, TypeRegion, Department, DepartmentNum, DistrictNum, MandateId) VALUES (?,?,?,?,?,?,?)"

		db := electionRepository.Sql.InitDB()

		stmt, err := db.Prepare((queryElection))
		if err != nil {
			electionRepository.Log.WriteErrorLog("Election Repository : Erreur préparation requête " + err.Error())
		}

		_, errExec := stmt.Exec(election.MandateCause, election.Region, election.TypeRegion, election.Department, election.DepartmentNum, election.DistrictNum, mandateId)
		if errExec != nil {
			electionRepository.Log.WriteErrorLog("Election Repository : Erreur exécution requête " + errExec.Error())
		}

		defer db.Close()
	}
}
