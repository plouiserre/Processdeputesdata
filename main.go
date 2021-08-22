package main


import (
	"fmt"
)

	//lire un fichier ici DID
	//le mettre dans un fichier à part
	//récupérer le contenu à afficher ensuite de l'objet
	//afficher le contenu

func main(){
	fmt.Println("Début lecture fichier")

	fileManager := fileManager {
		nameFile : "/Users/plouiserre/Projects/ProcessDeputesData/Data/1 - DeputesActifsMandatsActifsOrganes_XV/acteur/PA2960.json",
	}
	fileManager.getContentFile()

	fmt.Println("Contenu du fichier")
	fmt.Println(fileManager.contentFile)
	fmt.Println("Fin lecture fichier")
}