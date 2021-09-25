package main

type CongressManRepository struct {
	Sql  SqlManager
	Log  LogManager
	Data DataManager
}

func (repository *CongressManRepository) RecordAllCongressManData() int64 {
	congressManId := repository.RecordCongressManData()
	return congressManId
}

func (repository *CongressManRepository) RecordCongressManData() int64 {
	congressMan := repository.Data.CongressManModel
	queryCongressMan := "INSERT INTO PROCESSDEPUTES.Congressman(CongressManUid, Civility, FirstName, LastName, Alpha, Trigramme, BirthDate, BirthCity, BirthDepartment, BirthCountry, JobTitle, CatSocPro, FamSocPro) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"

	db := repository.Sql.InitDB()

	stmt, err := db.Prepare((queryCongressMan))
	if err != nil {
		repository.Log.WriteErrorLog("Erreur préparation requête " + err.Error())
	}

	//TODO supprimer la récupération de l'id après avoir ajouter l'ajout des mandates dans cette struct
	res, errExec := stmt.Exec(congressMan.CongressManUid, congressMan.Civility, congressMan.FirstName, congressMan.LastName, congressMan.Alpha, congressMan.Trigramme, congressMan.BirthDate, congressMan.BirthCity, congressMan.BirthDepartment, congressMan.BirthCountry, congressMan.JobTitle, congressMan.CatSocPro, congressMan.FamSocPro)
	if errExec != nil {
		repository.Log.WriteErrorLog("Erreur exécution requête " + errExec.Error())
	}

	lid, errGetLastId := res.LastInsertId()
	if errGetLastId != nil {
		repository.Log.WriteErrorLog("Erreur récupérage Id" + errGetLastId.Error())
	}

	defer db.Close()

	return lid
}
