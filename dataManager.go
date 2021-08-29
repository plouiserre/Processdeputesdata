package main

import (
	"encoding/json"
	"processdeputesdata/modelsJson"
	"log"
	"fmt"
)

type dataManager struct{
	CongressMan modelsJson.CongressMan
}

func (dataManager *dataManager) getDeputyData (congressManData string) {
	//fmt.Println(deputyData)
	data := []byte (congressManData)
	dataManager.deserializeData(data)
}

//TODO checker si on peut rendre cette méthode privée
func (dataManager *dataManager)deserializeData(data []byte){
	err := json.Unmarshal(data, &dataManager.CongressMan)
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
	fmt.Printf("Id du député %s \n", dataManager.CongressMan.Actor.Uid.Id)
}

func (dataManager *dataManager)displayIdendityDataDeserialize(){
	fmt.Printf("Civ du député %s \n",dataManager.CongressMan.Actor.CivilState.Identity.Civility)
	fmt.Printf("Prénom du député %s \n",dataManager.CongressMan.Actor.CivilState.Identity.FirstName)
	fmt.Printf("Nom du député %s \n",dataManager.CongressMan.Actor.CivilState.Identity.LastName)
	fmt.Printf("Alpha du député %s \n",dataManager.CongressMan.Actor.CivilState.Identity.Alpha)
	fmt.Printf("Trigramme du député %s \n",dataManager.CongressMan.Actor.CivilState.Identity.Trigramme)
}

func (dataManager *dataManager)displayBirthNewsDataDeserialize(){
	fmt.Printf("Date de naissance du député %s \n",dataManager.CongressMan.Actor.CivilState.BirthNews.BirthDate)
	fmt.Printf("Lieu de naissance du député%s \n",dataManager.CongressMan.Actor.CivilState.BirthNews.BirthCity)
	fmt.Printf("Département de naissance du député %s \n",dataManager.CongressMan.Actor.CivilState.BirthNews.BirthDepartment)
	fmt.Printf("Pays de naissance du député %s \n",dataManager.CongressMan.Actor.CivilState.BirthNews.BirthCountry)
}

func (dataManager *dataManager)displayProfessionDataDeserialize(){
	fmt.Printf("Libellé du travail du député %s \n",dataManager.CongressMan.Actor.Job.JobTitle)
	fmt.Printf("CatSocPro données de l'insee du travail du député %s \n",dataManager.CongressMan.Actor.Job.SocProcINSEE.CatSocPro)
	fmt.Printf("FamSocPro données de l'insee du travail du député %s \n",dataManager.CongressMan.Actor.Job.SocProcINSEE.FamSocPro)
}

func (dataManager *dataManager)displayAddressDataDeserialize(){
	for _ , address := range dataManager.CongressMan.Actor.Addresses.Address{
		if address.TypeAdress == "AdressePostale_Type"{
			fmt.Printf("Lieu %s \n %s \n %d %s \n %d %s \n", 
			address.TypeLabel, address.Label, address.StreetNumber, address.StreetName,
			address.PostalCode, address.City)
		}
	}
}

func (dataManager *dataManager)displayMandatsDataDeserialize(){
	for _, mandate := range dataManager.CongressMan.Actor.Mandates.Mandate {
		fmt.Printf("Id du mandat %s \nType d'organe %s \nNuméro de la Législature %d \nDate de début %s Date de fin %s\n", 
		mandate.Uid, mandate.TypeOrgane, mandate.TermOffice, mandate.StartDate, mandate.EndDate)	
		if mandate.Deputies != (modelsJson.Deputies{}){
			fmt.Printf("Id du suppléant %s, Date du Début %s et Date de fin %s\n", 
			mandate.Deputies.Deputy.RefDeputy, mandate.Deputies.Deputy.StartDate, mandate.Deputies.Deputy.EndDate)
		} 
		if mandate.Election != (modelsJson.Election{}){
			fmt.Printf("Cause de l'élection %s \nId de la circonscription %s\n", mandate.Election.MandateCause, mandate.Election.DistrictRef)
			fmt.Printf("Région %s, Type de Région de l'élection %s\nNom et Numéro  du département %s, %d, Numéro de la circonscription %d\n",
			mandate.Election.Place.Region, mandate.Election.Place.TypeRegion, 
			mandate.Election.Place.Department, mandate.Election.Place.DepartmentNum, mandate.Election.Place.DistrictNum)
		}
	}
}