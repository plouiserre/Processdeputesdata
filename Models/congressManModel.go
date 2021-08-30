package Models

type CongressManModel struct{
	CongressManUid string
	Civility string
	FirstName string
	LastName string
	Alpha string 
	Trigramme string
	BirthDate string
	BirthCity string
	BirthDepartment string
	BirthCountry string
	JobTitle string 
	CatSocPro string
	FamSocPro string
	Mandates []Mandate
}

func (congressManModel *CongressManModel) GetCongressManModel (congressManJson CongressManJson){
	congressManModel.CongressManUid = congressManJson.Actor.Uid.Id
	congressManModel.Civility = congressManJson.Actor.CivilState.Identity.Civility
	congressManModel.FirstName = congressManJson.Actor.CivilState.Identity.FirstName
	congressManModel.LastName = congressManJson.Actor.CivilState.Identity.LastName
	congressManModel.Alpha = congressManJson.Actor.CivilState.Identity.Alpha 
	congressManModel.Trigramme = congressManJson.Actor.CivilState.Identity.Trigramme
	congressManModel.BirthDate = congressManJson.Actor.CivilState.BirthNews.BirthDate
	congressManModel.BirthCity = congressManJson.Actor.CivilState.BirthNews.BirthCity
	congressManModel.BirthDepartment = congressManJson.Actor.CivilState.BirthNews.BirthDepartment
	congressManModel.BirthCountry = congressManJson.Actor.CivilState.BirthNews.BirthCountry
	congressManModel.JobTitle = congressManJson.Actor.Job.JobTitle 
	congressManModel.CatSocPro = congressManJson.Actor.Job.SocProcINSEE.CatSocPro
	congressManModel.FamSocPro = congressManJson.Actor.Job.SocProcINSEE.FamSocPro
	congressManModel.Mandates = []Mandate{}
	for _, mandate := range congressManJson.Actor.Mandates.Mandate{
		mandateModel := Mandate {}
		mandateModel.GetMandate(mandate)
		congressManModel.Mandates = append(congressManModel.Mandates,mandateModel)
		//congressManModel.Mandates[i].GetMandate(mandate) 
	}
}

type Mandate struct{
	MandateUid string 
	TermOffice int 
	TypeOrgane string
	StartDate string
	PublicationDate string
	EndDate string
	Precedence int
	PrincipleNomin int
	QualityCode string
	QualityLabel string
	QualityLabelSex string
	RefBody string
	Deputy Deputy
	Election Election
}
func (mandate *Mandate) GetMandate(mandateJson MandateJson){
	mandate.MandateUid = mandateJson.Uid
	mandate.TermOffice = mandateJson.TermOffice
	mandate.TypeOrgane = mandateJson.TypeOrgane
	mandate.StartDate = mandateJson.StartDate
	mandate.PublicationDate = mandateJson.PublicationDate
	mandate.EndDate = mandateJson.EndDate 
	mandate.Precedence = mandateJson.Precedence
	mandate.PrincipleNomin = mandateJson.PrincipleNomin
	mandate.QualityCode = mandateJson.DataQuality.QualityCode
	mandate.QualityLabel = mandateJson.DataQuality.QualityLabel
	mandate.QualityLabelSex = mandateJson.DataQuality.QualityLabelSex
	mandate.Deputy = Deputy{}
	mandate.Deputy.GetDeputy(mandateJson.Deputies)
	mandate.Election.GetElection(mandateJson.Election)
}

type Deputy struct{
	StartDate string `json:"dateDebut"`
	EndDate string `json:"dateFin"`
	RefDeputy string `json:"suppleantRef"`
}

func (deputy *Deputy) GetDeputy(deputiesJson DeputiesJson){
	deputy.RefDeputy = deputiesJson.Deputy.RefDeputy
	deputy.StartDate = deputiesJson.Deputy.StartDate
	deputy.EndDate = deputiesJson.Deputy.EndDate
}

type Election struct{
	MandateCause string `json:"causeMandat"`
	DistrictRef string `json:"refCirconscription"`
	Region string
	TypeRegion string `json:"regionType"`
	Department string `json:"departement"`
	DepartmentNum int `json:"numDepartement,string"`
	DistrictNum int `json:"numCirco,string"`
}

func (election *Election) GetElection(electionJson ElectionJson){
	election.MandateCause = electionJson.MandateCause
	election.DistrictRef = electionJson.DistrictRef
	election.Region = electionJson.Place.Region
	election.TypeRegion = electionJson.Place.TypeRegion
	election.Department = electionJson.Place.Department
	election.DepartmentNum = electionJson.Place.DepartmentNum
	election.DistrictNum = electionJson.Place.DistrictNum
}