package main

type RepositoryManager struct {
	Log           LogManager
	Data          DataManager
	Sql           SqlManager
	CongressManId int64
}

func (repositoryManager *RepositoryManager) StoreAllDatas() {
	repositoryManager.StoreCongressManDatas()
	repositoryManager.StoreMandateDatas()
	//repositoryManager.StoreDeputyDatas()
}

func (repositoryManager *RepositoryManager) StoreCongressManDatas() {
	congressManRepository := CongressManRepository{
		Log:  repositoryManager.Log,
		Sql:  repositoryManager.Sql,
		Data: repositoryManager.Data,
	}

	repositoryManager.CongressManId = congressManRepository.RecordAllCongressManData()
}

func (repositoryManager *RepositoryManager) StoreMandateDatas() {
	mandateRepository := MandateRepository{
		Log:  repositoryManager.Log,
		Sql:  repositoryManager.Sql,
		Data: repositoryManager.Data,
	}

	mandateRepository.RecordAllMandates(repositoryManager.CongressManId)
}

/*func (repositoryManager *RepositoryManager) StoreDeputyDatas() {
	deputyRepository := DeputyRepository{
		Log:  repositoryManager.Log,
		Sql:  repositoryManager.Sql,
		Data: repositoryManager.Data,
	}

	deputyRepository.RecordAllDeputyData()
}*/
