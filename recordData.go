package main

//TODO faut tout mettre dans des classes spécifiques pour chaque entités et renommer cette classe

type RecordData struct {
	Sql  SqlManager
	Log  LogManager
	Data DataManager
}

func (recordData *RecordData) RecordAllCongressManData() {
	recordData.RecordCongressManData()
}

func (recordData *RecordData) RecordCongressManData() {
	congressMan := recordData.Data.CongressManModel
	queryCongressMan := "INSERT INTO PROCESSDEPUTES.Congressman(CongressManUid, Civility, FirstName, LastName, Alpha, Trigramme, BirthDate, BirthCity, BirthDepartment, BirthCountry, JobTitle, CatSocPro, FamSocPro) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"
	//recordData.Sql.InsertData(queryCongressMan, nil, congressMan.CongressManUid, congressMan.Civility, congressMan.FirstName, congressMan.LastName, congressMan.Alpha, congressMan.Trigramme, congressMan.BirthDate, congressMan.BirthCity, congressMan.BirthDepartment, congressMan.BirthCountry, congressMan.JobTitle, congressMan.CatSocPro, congressMan.FamSocPro)

	db := recordData.Sql.InitDB()

	stmt, err := db.Prepare((queryCongressMan))
	if err != nil {
		recordData.Log.WriteErrorLog("Erreur préparation requête " + err.Error())
	}

	//recordData.Log.WriteInfoLog("Test " + congressMan.CongressManUid + " " + congressMan.Civility + " " + congressMan.FirstName)

	_, errExec := stmt.Exec(congressMan.CongressManUid, congressMan.Civility, congressMan.FirstName, congressMan.LastName, congressMan.Alpha, congressMan.Trigramme, congressMan.BirthDate, congressMan.BirthCity, congressMan.BirthDepartment, congressMan.BirthCountry, congressMan.JobTitle, congressMan.CatSocPro, congressMan.FamSocPro)
	if errExec != nil {
		recordData.Log.WriteErrorLog("Erreur exécution requête " + errExec.Error())
	}

	defer db.Close()
}

/*func (sqlManager *SqlManager) InsertData(insertQuery string, args ...interface{}) {
	db := sqlManager.InitDB()
	stmt, err := db.Prepare((insertQuery))
	if err != nil {
		sqlManager.Log.WriteErrorLog("Erreur préparation requête " + err.Error())
	}
	insertArgs := []interface{}{args}
	_, errExec := stmt.Exec(insertArgs)
	if errExec != nil {
		sqlManager.Log.WriteErrorLog("Erreur exécution requête " + errExec.Error())
	}
	//sqlManager.Log.WriteInfoLog("Request " + logQuery)
	//TODO améliorer cette ligne
	//sqlManager.Log.WriteInfoLog("Request " + logQuery)
	defer db.Close()
}*/
