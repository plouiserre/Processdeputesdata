package main


func main(){
	logManager := LogManager{}
	logManager.InitLog()

	logManager.WriteInfoLog("Début lecture fichier")

	fileManager := fileManager {
		nameFile : "/Users/plouiserre/Projects/ProcessDeputesData/Data/1 - DeputesActifsMandatsActifsOrganes_XV/acteur/PA2960.json",
		LogManager : logManager,
	}
	fileManager.getContentFile()

	logManager.WriteInfoLog("Début Désérialisation")
	dataManager := dataManager {
		LogManager : logManager,
	}
	dataManager.ProcessDeputyData(fileManager.contentFile)
	logManager.WriteInfoLog("Fin désérialisation")
}