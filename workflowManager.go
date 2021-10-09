package main

import (
	"io/ioutil"
	"time"
)

type WorkflowManager struct {
	LogManager LogManager
	FolderName string
}

func (workflowManager *WorkflowManager) StoreAllDatas() {
	startWorkflow := time.Now()
	files, _ := ioutil.ReadDir(workflowManager.FolderName)

	for _, file := range files {
		workflowManager.StoreDatasFile(file.Name())
	}
	elapsedWorkflow := time.Since(startWorkflow)
	workflowManager.LogManager.WriteInfoLog("Le workflow a été exécuté en " + elapsedWorkflow.String())
}

func (workflowManager *WorkflowManager) StoreDatasFile(nameFile string) {
	startFile := time.Now()
	workflowManager.LogManager.WriteInfoLog("Début Traitement Fichier " + nameFile)

	fileManager := workflowManager.GetFile(nameFile)

	dataManager := workflowManager.DeserializeFile(fileManager)

	workflowManager.StoreDatas(dataManager)

	elapsedFile := time.Since(startFile)
	workflowManager.LogManager.WriteInfoLog("Fichier " + nameFile + " a été traité en " + elapsedFile.String())
}

func (workflowManager *WorkflowManager) GetFile(nameFile string) fileManager {
	locationFileComplete := workflowManager.FolderName + nameFile
	fileManager := fileManager{
		nameFile:   locationFileComplete,
		LogManager: workflowManager.LogManager,
	}
	fileManager.getContentFile()

	return fileManager
}

func (workflowManager *WorkflowManager) DeserializeFile(fileManager fileManager) DataManager {
	startDeserialize := time.Now()
	workflowManager.LogManager.WriteInfoLog("Début Désérialisation")
	dataManager := DataManager{
		LogManager: workflowManager.LogManager,
	}
	dataManager.ProcessDeputyData(fileManager.contentFile)
	elapsedDeserialize := time.Since(startDeserialize)
	workflowManager.LogManager.WriteInfoLog("Fin désérialisation en " + elapsedDeserialize.String())

	return dataManager
}

func (workflowManager *WorkflowManager) StoreDatas(dataManager DataManager) {
	startStoreDatas := time.Now()
	workflowManager.LogManager.WriteInfoLog("Début enregistrement")

	repositoryManager := RepositoryManager{
		Log:  workflowManager.LogManager,
		Data: dataManager,
	}
	repositoryManager.StoreAllDatas()
	elapsedStoreDatas := time.Since(startStoreDatas)
	workflowManager.LogManager.WriteInfoLog("Fin enregistrement en " + elapsedStoreDatas.String())
}
