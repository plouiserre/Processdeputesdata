package main

type ElectionRepository struct {
	RepositoryManager *RepositoryManager
}

func (electionRepository *ElectionRepository) RecordElection(mandateId int64, election ElectionModel) {
	if election != (ElectionModel{}) {
		repository := electionRepository.RepositoryManager
		queryElection := "INSERT INTO PROCESSDEPUTES.Election(MandateCause, Region, TypeRegion, Department, DepartmentNum, DistrictNum, MandateId) VALUES (?,?,?,?,?,?,?)"
		nameRepository := "election Repository"

		db := repository.Sql.InitDB()

		stmt, isOk := repository.Sql.PrepareRequest(db, queryElection, nameRepository)

		if isOk {
			_, errExec := stmt.Exec(election.MandateCause, election.Region, election.TypeRegion, election.Department, election.DepartmentNum, election.DistrictNum, mandateId)
			if errExec != nil {
				repository.Log.WriteErrorLog("Election Repository : Erreur exécution requête " + errExec.Error())
			}
		}

		defer db.Close()
	}
}
