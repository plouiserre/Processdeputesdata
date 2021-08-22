package main


import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
)

	//lire un fichier ici DID
	//le mettre dans un fichier à part
	//récupérer le contenu à afficher ensuite de l'objet
	//afficher le contenu

func main(){
	fmt.Println("Début lecture fichier")

	file, err := os.OpenFile("/Users/plouiserre/Projects/ProcessDeputesData/Data/1 - DeputesActifsMandatsActifsOrganes_XV/acteur/PA2960.json",  os.O_RDONLY,0666)
	if err != nil{
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil{
		log.Fatal(err)
	}
	displayContnet := string(content)
	fmt.Println("Contenu du fichier")
	fmt.Print(displayContnet)
	fmt.Println("Fin lecture fichier")
}