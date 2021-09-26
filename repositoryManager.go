package main

type RepositoryManager struct {
	Log  LogManager
	Data DataManager
	Sql  SqlManager
}

func (repositoryManager *RepositoryManager) StoreAllDatas() {
	repositoryManager.StoreCongressManDatas()
}

func (repositoryManager *RepositoryManager) StoreCongressManDatas() {
	deputyRepository := DeputyRepository{
		Log:  repositoryManager.Log,
		Sql:  repositoryManager.Sql,
		Data: repositoryManager.Data,
	}
	electionRepository := ElectionRepository{
		Log:  repositoryManager.Log,
		Sql:  repositoryManager.Sql,
		Data: repositoryManager.Data,
	}
	mandateRepository := MandateRepository{
		Log:                repositoryManager.Log,
		Sql:                repositoryManager.Sql,
		Data:               repositoryManager.Data,
		DeputyRepository:   deputyRepository,
		ElectionRepository: electionRepository,
	}
	congressManRepository := CongressManRepository{
		Log:               repositoryManager.Log,
		Sql:               repositoryManager.Sql,
		Data:              repositoryManager.Data,
		MandateRepository: mandateRepository,
	}

	congressManRepository.RecordAllCongressManData()
}
