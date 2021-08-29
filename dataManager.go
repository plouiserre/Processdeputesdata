package main

//TODO externaliser la partie log plus tard dans un nouveau fichier pour la partie workflow

import (
	"encoding/json"
	"processdeputesdata/Models"
	"log"
	"fmt"
)

type dataManager struct{
	CongressManJson Models.CongressManJson
	CongressManModel Models.CongressManModel
}

func (dataManager *dataManager) ProcessDeputyData(congressManData string) {
	dataManager.GetDeputyDataJson(congressManData)
	
	dataManager.CongressManModel.GetCongressManModel(dataManager.CongressManJson)
	
	dataManager.DisplayCongressManModel()
}

//TODO checker si on peut rendre cette méthode privée
func (dataManager *dataManager)GetDeputyDataJson(congressManData string){
	data := []byte (congressManData)
	err := json.Unmarshal(data, &dataManager.CongressManJson)
	if err != nil{
		log.Fatal(err)
	}
	dataManager.displayDataDeserialize()
}

//TODO à supprimer
func (dataManager *dataManager)displayDataDeserialize(){
	dataManager.displayUidDataDeserialize()
	dataManager.displayIdendityDataDeserialize()
	dataManager.displayBirthNewsDataDeserialize()
	dataManager.displayProfessionDataDeserialize()
	dataManager.displayAddressDataDeserialize()
	dataManager.displayMandatsDataDeserialize()
}

func (dataManager *dataManager)displayUidDataDeserialize(){
	fmt.Printf("Id du député %s \n", dataManager.CongressManJson.Actor.Uid.Id)
}

func (dataManager *dataManager)displayIdendityDataDeserialize(){
	fmt.Printf("Civ du député %s \n",dataManager.CongressManJson.Actor.CivilState.Identity.Civility)
	fmt.Printf("Prénom du député %s \n",dataManager.CongressManJson.Actor.CivilState.Identity.FirstName)
	fmt.Printf("Nom du député %s \n",dataManager.CongressManJson.Actor.CivilState.Identity.LastName)
	fmt.Printf("Alpha du député %s \n",dataManager.CongressManJson.Actor.CivilState.Identity.Alpha)
	fmt.Printf("Trigramme du député %s \n",dataManager.CongressManJson.Actor.CivilState.Identity.Trigramme)
}

func (dataManager *dataManager)displayBirthNewsDataDeserialize(){
	fmt.Printf("Date de naissance du député %s \n",dataManager.CongressManJson.Actor.CivilState.BirthNews.BirthDate)
	fmt.Printf("Lieu de naissance du député%s \n",dataManager.CongressManJson.Actor.CivilState.BirthNews.BirthCity)
	fmt.Printf("Département de naissance du député %s \n",dataManager.CongressManJson.Actor.CivilState.BirthNews.BirthDepartment)
	fmt.Printf("Pays de naissance du député %s \n",dataManager.CongressManJson.Actor.CivilState.BirthNews.BirthCountry)
}

func (dataManager *dataManager)displayProfessionDataDeserialize(){
	fmt.Printf("Libellé du travail du député %s \n",dataManager.CongressManJson.Actor.Job.JobTitle)
	fmt.Printf("CatSocPro données de l'insee du travail du député %s \n",dataManager.CongressManJson.Actor.Job.SocProcINSEE.CatSocPro)
	fmt.Printf("FamSocPro données de l'insee du travail du député %s \n",dataManager.CongressManJson.Actor.Job.SocProcINSEE.FamSocPro)
}

func (dataManager *dataManager)displayAddressDataDeserialize(){
	for _ , address := range dataManager.CongressManJson.Actor.Addresses.Address{
		if address.TypeAdress == "AdressePostale_Type"{
			fmt.Printf("Lieu %s \n %s \n %d %s \n %d %s \n", 
			address.TypeLabel, address.Label, address.StreetNumber, address.StreetName,
			address.PostalCode, address.City)
		}
	}
}

func (dataManager *dataManager)displayMandatsDataDeserialize(){
	for _, mandate := range dataManager.CongressManJson.Actor.Mandates.Mandate {
		fmt.Printf("Id du mandat %s \nType d'organe %s \nNuméro de la Législature %d \nDate de début %s Date de fin %s\n", 
		mandate.Uid, mandate.TypeOrgane, mandate.TermOffice, mandate.StartDate, mandate.EndDate)	
		if mandate.Deputies != (Models.DeputiesJson{}){
			fmt.Printf("Id du suppléant %s, Date du Début %s et Date de fin %s\n", 
			mandate.Deputies.Deputy.RefDeputy, mandate.Deputies.Deputy.StartDate, mandate.Deputies.Deputy.EndDate)
		} 
		if mandate.Election != (Models.ElectionJson{}){
			fmt.Printf("Cause de l'élection %s \nId de la circonscription %s\n", mandate.Election.MandateCause, mandate.Election.DistrictRef)
			fmt.Printf("Région %s, Type de Région de l'élection %s\nNom et Numéro  du département %s, %d, Numéro de la circonscription %d\n",
			mandate.Election.Place.Region, mandate.Election.Place.TypeRegion, 
			mandate.Election.Place.Department, mandate.Election.Place.DepartmentNum, mandate.Election.Place.DistrictNum)
		}
	}
}

func (dataManager *dataManager) DisplayCongressManModel(){
	fmt.Println("Début affichage donnée CongressMan Model ")
	fmt.Printf("Id %s\n",dataManager.CongressManModel.CongressManUid)
	fmt.Printf("Civility %s\n",dataManager.CongressManModel.Civility)
	fmt.Printf("FirstName %s\n",dataManager.CongressManModel.FirstName)
	fmt.Printf("LastName %s\n",dataManager.CongressManModel.LastName)
	fmt.Printf("Alpha %s\n",dataManager.CongressManModel.Alpha)
	fmt.Printf("Trigramme %s\n",dataManager.CongressManModel.Trigramme)
	fmt.Printf("BirthDate %s\n",dataManager.CongressManModel.BirthDate)
	fmt.Printf("BirthCity %s\n",dataManager.CongressManModel.BirthCity)
	fmt.Printf("BirthDepartment %s\n",dataManager.CongressManModel.BirthDepartment)
	fmt.Printf("BirthCountry %s\n",dataManager.CongressManModel.BirthCountry)
	fmt.Printf("JobTitle %s\n",dataManager.CongressManModel.JobTitle)
	fmt.Printf("CatSocPro %s\n",dataManager.CongressManModel.CatSocPro)
	fmt.Printf("FamSocPro %s\n",dataManager.CongressManModel.FamSocPro)
	fmt.Printf("BirthDepartment %s\n",dataManager.CongressManModel.BirthDepartment)
	
	dataManager.DisplayMandatesModel()

	fmt.Println("Fin affichage donnée CongressMan Model ")
}

func (dataManager *dataManager) DisplayMandatesModel(){
	for _, mandate := range dataManager.CongressManModel.Mandates {
		fmt.Printf("Id du mandat %s \nType d'organe %s \nNuméro de la Législature %d \nDate de début %s Date de fin %s\n", 
		mandate.MandateUid, mandate.TypeOrgane, mandate.TermOffice, mandate.StartDate, mandate.EndDate)	
		if mandate.Deputy != (Models.Deputy{}){
			fmt.Printf("Id du suppléant %s, Date du Début %s et Date de fin %s\n", 
			mandate.Deputy.RefDeputy, mandate.Deputy.StartDate, mandate.Deputy.EndDate)
		} 
		if mandate.Election != (Models.Election{}){
			fmt.Printf("Cause de l'élection %s \nId de la circonscription %s\n", mandate.Election.MandateCause, mandate.Election.DistrictRef)
			fmt.Printf("Région %s, Type de Région de l'élection %s\nNom et Numéro  du département %s, %d, Numéro de la circonscription %d\n",
			mandate.Election.Region, mandate.Election.TypeRegion, mandate.Election.Department, mandate.Election.DepartmentNum, mandate.Election.DistrictNum)
		}
	}
}
