package main

func main() {
	logManager := LogManager{}
	logManager.InitLog()

	logManager.WriteInfoLog("Début lecture fichier")
	workflowManager := WorkflowManager{
		LogManager: logManager,
		FolderName: "/Users/plouiserre/Projects/ProcessDeputesData/Data/1 - DeputesActifsMandatsActifsOrganes_XV/acteur/",
	}

	workflowManager.StoreAllDatas()
}
