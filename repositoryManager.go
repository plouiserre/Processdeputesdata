package main

type RepositoryManager struct {
	Log  LogManager
	Data DataManager
	Sql  SqlManager
}

func (repositoryManager *RepositoryManager) StoreAllDatas() {
	repositoryManager.StoreCongressManDatas()
	repositoryManager.StoreDeputyDatas()
}

func (repositoryManager *RepositoryManager) StoreCongressManDatas() {
	congressManRepository := CongressManRepository{
		Log:  repositoryManager.Log,
		Sql:  repositoryManager.Sql,
		Data: repositoryManager.Data,
	}

	congressManRepository.RecordAllCongressManData()
}

func (repositoryManager *RepositoryManager) StoreDeputyDatas() {
	deputyRepository := DeputyRepository{
		Log:  repositoryManager.Log,
		Sql:  repositoryManager.Sql,
		Data: repositoryManager.Data,
	}

	deputyRepository.RecordAllDeputyData()
}
