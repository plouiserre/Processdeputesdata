package modelsJson

//all struct get in 1 - DeputesActifsMandatsActifsOrganes_XV

type CongressMan struct{
	Actor Actor `json:"Acteur"`
}

type Actor struct{
	Uid Uid
	CivilState CivilState `json:"etatCivil"`
	Job Profession `json:"Profession"`
	Addresses Addresses `json:"Adresses"`
	Mandates Mandates `json:"Mandats"` 
} 

type Uid struct{
	Id string `json:"#text"`
}

type CivilState struct{
	Identity Identity `json:"ident"`
	BirthNews BirthNews	 `json:"infoNaissance"`
}

type Identity struct{
	Civility string `json:"civ"`
	FirstName string `json:"prenom"`
	LastName string `json:"nom"`
	Alpha string 
	Trigramme string
}

type BirthNews struct{
	BirthDate string `json:"dateNais"`
	BirthCity string `json:"villeNais"`
	BirthDepartment string `json:"depNais"`
	BirthCountry string `json:"paysNais"`
}

type Profession struct{
	JobTitle string `json:"libelleCourant"`
	SocProcINSEE JobInseeData
}

type JobInseeData struct{
	CatSocPro string
	FamSocPro string
}

type Addresses struct{
	Address []Address `json:"Adresse"`
}

type Address struct{
	uid string
	typeId int //Warning ici aussi
	TypeAdress string `json:"@xsi:type"`
	TypeLabel string `json:"TypeLibelle"`
	poids int 
	AdresseDeRattachement string
	Label string `json:"Intitule"`
	StreetNumber int `json:"NumeroRue,string"`
	StreetName string `json:"NomRue"`
	ComplementAdresse string
	PostalCode int `json:"CodePostal,string"`
	City string `json:"Ville"`
	valElect string
}

type Mandates struct{
	Mandate []Mandate `json:"Mandat"`
}

type Mandate struct{
	Uid string 
	ActeurId string `json:"ActeurRef"`
	TermOffice int `json:"Legislature,string"`
	TypeOrgane string
	StartDate string `json:"dateDebut"`
	PublicationDate string `json:"datePublication"`
	EndDate string `json:"dateFin"`
	Precedence int `json:"preseance,string"`
	PrincipleNomin int `json:"nominPrincipale,string"`
	DataQuality DataQuality `json:"infosQualite"`
	Body Body `json:"organe"`
	Deputies Deputies `json:"suppleants"`
	Election Election
}

type DataQuality struct{
	QualityCode string `json:"codeQualite"`
	QualiteLabel string `json:"libQualite"`
	QualiteLabelSex string `json:"libQualiteSex"`
}

type Body struct{
	RefBody string `json:"organeRef"`
}

type Deputies struct{
	Deputy Deputy `json:"suppleant"`
}

type Deputy struct{
	StartDate string `json:"dateDebut"`
	EndDate string `json:"dateFin"`
	RefDeputy string `json:"suppleantRef"`
}

type Election struct{
	MandateCause string `json:"causeMandat"`
	DistrictRef string `json:"refCirconscription"`
	Place Place `json:"lieu"`
}

type Place struct{
	Region string
	TypeRegion string `json:"regionType"`
	Department string `json:"departement"`
	DepartmentNum int `json:"numDepartement,string"`
	DistrictNum int `json:"numCirco,string"`
}