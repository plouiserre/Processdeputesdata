package main

import "io/ioutil"

type WorkflowManager struct {
	LogManager LogManager
	FolderName string
}

func (workflowManager *WorkflowManager) StoreAllDatas() {
	files, _ := ioutil.ReadDir(workflowManager.FolderName)

	for _, file := range files {
		workflowManager.StoreDatasFile(file.Name())
	}
}

func (workflowManager *WorkflowManager) StoreDatasFile(nameFile string) {
	locationFileComplete := workflowManager.FolderName + nameFile
	workflowManager.LogManager.WriteInfoLog("Début Traitement Fichier " + nameFile)
	fileManager := fileManager{
		nameFile:   locationFileComplete,
		LogManager: workflowManager.LogManager,
	}
	fileManager.getContentFile()

	workflowManager.LogManager.WriteInfoLog("Début Désérialisation")
	dataManager := DataManager{
		LogManager: workflowManager.LogManager,
	}
	dataManager.ProcessDeputyData(fileManager.contentFile)
	workflowManager.LogManager.WriteInfoLog("Fin désérialisation")

	workflowManager.LogManager.WriteInfoLog("Début enregistrement")

	repositoryManager := RepositoryManager{
		Log:  workflowManager.LogManager,
		Data: dataManager,
	}
	repositoryManager.StoreAllDatas()
	workflowManager.LogManager.WriteInfoLog("Fin enregistrement")

	workflowManager.LogManager.WriteInfoLog("Fin Traitement Fichier " + nameFile)
}
