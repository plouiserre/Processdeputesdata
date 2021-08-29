package main


import (
	"fmt"
)

func main(){
	fmt.Println("Début lecture fichier")

	fileManager := fileManager {
		nameFile : "/Users/plouiserre/Projects/ProcessDeputesData/Data/1 - DeputesActifsMandatsActifsOrganes_XV/acteur/PA2960.json",
	}
	fileManager.getContentFile()

	/*fmt.Println("Contenu du fichier")
	fmt.Println(fileManager.contentFile)
	fmt.Println("Fin lecture fichier")*/
	fmt.Println("Début Désérialisation")
	dataManager := dataManager {
	}
	dataManager.getDeputyData(fileManager.contentFile)
	fmt.Println("Fin désérialisation")
}