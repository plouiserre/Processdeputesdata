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
		RepositoryManager: repositoryManager,
	}
	electionRepository := ElectionRepository{
		RepositoryManager: repositoryManager,
	}
	mandateRepository := MandateRepository{
		RepositoryManager:  repositoryManager,
		DeputyRepository:   deputyRepository,
		ElectionRepository: electionRepository,
	}
	congressManRepository := CongressManRepository{
		RepositoryManager: repositoryManager,
		MandateRepository: mandateRepository,
	}

	congressManRepository.RecordAllCongressManData()
}
