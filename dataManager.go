package main

import (
	"encoding/json"
	"log"
)

type DataManager struct {
	CongressManJson  CongressManJson
	CongressManModel CongressManModel
	LogManager       LogManager
}

func (dataManager *DataManager) ProcessDeputyData(congressManData string) {
	dataManager.getDeputyDataJson(congressManData)

	dataManager.CongressManModel.GetCongressManModel(dataManager.CongressManJson)
}

func (dataManager *DataManager) getDeputyDataJson(congressManData string) {
	data := []byte(congressManData)
	err := json.Unmarshal(data, &dataManager.CongressManJson)
	if err != nil {
		log.Fatal(err)
	}
}
