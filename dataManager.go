package main

//TODO externaliser la partie log plus tard dans un nouveau fichier pour la partie workflow

import (
	"encoding/json"
	"log"
	"processdeputesdata/Models/Models"
	//"processdeputesdata/Models"
)

type DataManager struct {
	CongressManJson  Models.CongressManJson
	CongressManModel Models.CongressManModel
	LogManager       LogManager
}

func (dataManager *DataManager) ProcessDeputyData(congressManData string) {
	dataManager.GetDeputyDataJson(congressManData)

	dataManager.CongressManModel.GetCongressManModel(dataManager.CongressManJson)

	//dataManager.DisplayCongressManModel()
}

//TODO checker si on peut rendre cette méthode privée
func (dataManager *DataManager) GetDeputyDataJson(congressManData string) {
	//dataManager.LogManager.WriteInfoLog(congressManData)
	data := []byte(congressManData)
	err := json.Unmarshal(data, &dataManager.CongressManJson)
	if err != nil {
		log.Fatal(err)
	}
	//dataManager.displayDataDeserialize()
}

//TODO à supprimer
/*
func (dataManager *DataManager) displayDataDeserialize() {
	dataManager.displayUidDataDeserialize()
	dataManager.displayIdendityDataDeserialize()
	dataManager.displayBirthNewsDataDeserialize()
	dataManager.displayProfessionDataDeserialize()
	dataManager.displayAddressDataDeserialize()
	dataManager.displayMandatsDataDeserialize()
}

func (dataManager *DataManager) displayUidDataDeserialize() {
	dataManager.LogManager.WriteInfoLog("Id du député " + dataManager.CongressManJson.Actor.Uid.Id)
}

func (dataManager *DataManager) displayIdendityDataDeserialize() {
	dataManager.LogManager.WriteInfoLog("Civ du député " + dataManager.CongressManJson.Actor.CivilState.Identity.Civility)
	dataManager.LogManager.WriteInfoLog("Prénom du député " + dataManager.CongressManJson.Actor.CivilState.Identity.FirstName)
	dataManager.LogManager.WriteInfoLog("Nom du député " + dataManager.CongressManJson.Actor.CivilState.Identity.LastName)
	dataManager.LogManager.WriteInfoLog("Alpha du député " + dataManager.CongressManJson.Actor.CivilState.Identity.Alpha)
	dataManager.LogManager.WriteInfoLog("Trigramme du député " + dataManager.CongressManJson.Actor.CivilState.Identity.Trigramme)
}

func (dataManager *DataManager) displayBirthNewsDataDeserialize() {
	dataManager.LogManager.WriteInfoLog("Date de naissance du député " + dataManager.CongressManJson.Actor.CivilState.BirthNews.BirthDate)
	dataManager.LogManager.WriteInfoLog("Lieu de naissance du député " + dataManager.CongressManJson.Actor.CivilState.BirthNews.BirthCity)
	dataManager.LogManager.WriteInfoLog("Département de naissance du député " + dataManager.CongressManJson.Actor.CivilState.BirthNews.BirthDepartment)
	dataManager.LogManager.WriteInfoLog("Pays de naissance du député " + dataManager.CongressManJson.Actor.CivilState.BirthNews.BirthCountry)
}

func (dataManager *DataManager) displayProfessionDataDeserialize() {
	dataManager.LogManager.WriteInfoLog("Libellé du travail du député " + dataManager.CongressManJson.Actor.Job.JobTitle)
	dataManager.LogManager.WriteInfoLog("CatSocPro données de l'insee du travail du député " + dataManager.CongressManJson.Actor.Job.SocProcINSEE.CatSocPro)
	dataManager.LogManager.WriteInfoLog("FamSocPro données de l'insee du travail du député " + dataManager.CongressManJson.Actor.Job.SocProcINSEE.FamSocPro)
}

func (dataManager *DataManager) displayAddressDataDeserialize() {
	for _, address := range dataManager.CongressManJson.Actor.Addresses.Address {
		if address.TypeAdress == "AdressePostale_Type" {
			dataManager.LogManager.WriteInfoLog("Lieu " + address.TypeLabel)
			dataManager.LogManager.WriteInfoLog(address.Label)
			dataManager.LogManager.WriteInfoLog(string(address.StreetNumber) + " " + address.StreetName)
			dataManager.LogManager.WriteInfoLog(string(address.PostalCode) + " " + address.City)
		}
	}
}

func (dataManager *DataManager) displayMandatsDataDeserialize() {
	for _, mandate := range dataManager.CongressManJson.Actor.Mandates.Mandate {
		dataManager.LogManager.WriteInfoLog("Id du mandat " + mandate.Uid)
		dataManager.LogManager.WriteInfoLog("Type d'organe " + mandate.TypeOrgane)
		dataManager.LogManager.WriteInfoLog("Numéro de la Législature " + string(mandate.TermOffice))
		dataManager.LogManager.WriteInfoLog("Date de début " + mandate.StartDate + " Date de fin " + mandate.EndDate)
		if mandate.Deputies != (Models.DeputiesJson{}) {
			dataManager.LogManager.WriteInfoLog("Id du suppléant " + mandate.Deputies.Deputy.RefDeputy + " Date du Début " +
				mandate.Deputies.Deputy.StartDate + " et Date de fin " + mandate.Deputies.Deputy.EndDate)
		}
		if mandate.Election != (Models.ElectionJson{}) {
			dataManager.LogManager.WriteInfoLog("Cause de l'élection " + mandate.Election.MandateCause)
			dataManager.LogManager.WriteInfoLog("Id de la circonscription " + mandate.Election.DistrictRef)
			dataManager.LogManager.WriteInfoLog("Région " + mandate.Election.Place.Region + " Type de Région de l'élection " + mandate.Election.Place.TypeRegion)
			dataManager.LogManager.WriteInfoLog("Nom et Numéro du département " + mandate.Election.Place.Department + " " + string(mandate.Election.Place.DepartmentNum) +
				"Numéro de la circonscription " + string(mandate.Election.Place.DistrictNum))
		}
	}
}

func (dataManager *DataManager) DisplayCongressManModel() {
	dataManager.LogManager.WriteInfoLog("Début affichage donnée CongressMan Model ")
	dataManager.LogManager.WriteInfoLog("Id " + dataManager.CongressManModel.CongressManUid)
	dataManager.LogManager.WriteInfoLog("Civility " + dataManager.CongressManModel.Civility)
	dataManager.LogManager.WriteInfoLog("FirstName " + dataManager.CongressManModel.FirstName)
	dataManager.LogManager.WriteInfoLog("LastName  " + dataManager.CongressManModel.LastName)
	dataManager.LogManager.WriteInfoLog("Alpha " + dataManager.CongressManModel.Alpha)
	dataManager.LogManager.WriteInfoLog("Trigramme " + dataManager.CongressManModel.Trigramme)
	dataManager.LogManager.WriteInfoLog("BirthDate " + dataManager.CongressManModel.BirthDate)
	dataManager.LogManager.WriteInfoLog("BirthCity " + dataManager.CongressManModel.BirthCity)
	dataManager.LogManager.WriteInfoLog("BirthDepartment " + dataManager.CongressManModel.BirthDepartment)
	dataManager.LogManager.WriteInfoLog("BirthCountry " + dataManager.CongressManModel.BirthCountry)
	dataManager.LogManager.WriteInfoLog("JobTitle " + dataManager.CongressManModel.JobTitle)
	dataManager.LogManager.WriteInfoLog("CatSocPro " + dataManager.CongressManModel.CatSocPro)
	dataManager.LogManager.WriteInfoLog("FamSocPro " + dataManager.CongressManModel.FamSocPro)
	dataManager.LogManager.WriteInfoLog("BirthDepartment " + dataManager.CongressManModel.BirthDepartment)

	dataManager.DisplayMandatesModel()

	dataManager.LogManager.WriteInfoLog("Fin affichage donnée CongressMan Model ")
}

func (dataManager *DataManager) DisplayMandatesModel() {
	for _, mandate := range dataManager.CongressManModel.Mandates {
		dataManager.LogManager.WriteInfoLog("Id du mandat " + mandate.MandateUid)
		dataManager.LogManager.WriteInfoLog("Type d'organe " + mandate.TypeOrgane)
		dataManager.LogManager.WriteInfoLog("Numéro de la Législature " + string(mandate.TermOffice))
		dataManager.LogManager.WriteInfoLog("Date du début " + mandate.StartDate + " Date de fin " + mandate.EndDate)
		if mandate.Deputy != (Models.Deputy{}) {
			dataManager.LogManager.WriteInfoLog("Id du suppléant " + mandate.Deputy.RefDeputy + " Date du début " + mandate.Deputy.StartDate + " Date de fin " + mandate.Deputy.EndDate)
		}
		if mandate.Election != (Models.Election{}) {
			dataManager.LogManager.WriteInfoLog("Cause de l'élection " + mandate.Election.MandateCause)
			dataManager.LogManager.WriteInfoLog("Id de la circonscription " + mandate.Election.DistrictRef)
			dataManager.LogManager.WriteInfoLog("Région " + mandate.Election.Region + " Type de Région de l'élection " + mandate.Election.TypeRegion)
			dataManager.LogManager.WriteInfoLog("Nom et Numéro  du département " + mandate.Election.Department + " " + string(mandate.Election.DepartmentNum) + " Numéro de la circonscription " + string(mandate.Election.DistrictNum))
		}
	}
}
*/
