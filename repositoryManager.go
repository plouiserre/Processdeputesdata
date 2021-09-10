package main

type RepositoryManager struct {
	Log  LogManager
	Data DataManager
}

func (repositoryManager *RepositoryManager) StoreAllDatas() {
	sqlManager := SqlManager{
		Log: repositoryManager.Log,
	}
	repository := CongressManRepository{
		Log:  repositoryManager.Log,
		Sql:  sqlManager,
		Data: repositoryManager.Data,
	}

	repository.RecordAllCongressManData()
}
