package main

//all struct get in 1 - DeputesActifsMandatsActifsOrganes_XV

type CongressManJson struct {
	Actor ActorJson `json:"Acteur"`
}

type ActorJson struct {
	Uid        UidJson
	CivilState CivilStateJson `json:"etatCivil"`
	Job        ProfessionJson `json:"Profession"`
	Addresses  AddressesJson  `json:"Adresses"`
	Mandates   MandatesJson   `json:"Mandats"`
}

type UidJson struct {
	Id string `json:"#text"`
}

type CivilStateJson struct {
	Identity  IdentityJson  `json:"ident"`
	BirthNews BirthNewsJson `json:"infoNaissance"`
}

type IdentityJson struct {
	Civility  string `json:"civ"`
	FirstName string `json:"prenom"`
	LastName  string `json:"nom"`
	Alpha     string
	Trigramme string
}

type BirthNewsJson struct {
	BirthDate       string `json:"dateNais"`
	BirthCity       string `json:"villeNais"`
	BirthDepartment string `json:"depNais"`
	BirthCountry    string `json:"paysNais"`
}

type ProfessionJson struct {
	JobTitle     string `json:"libelleCourant"`
	SocProcINSEE JobInseeDataJson
}

type JobInseeDataJson struct {
	CatSocPro string
	FamSocPro string
}

type AddressesJson struct {
	Address []Address `json:"Adresse"`
}

type Address struct {
	uid                   string
	typeId                int    //Warning ici aussi
	TypeAdress            string `json:"@xsi:type"`
	TypeLabel             string `json:"TypeLibelle"`
	poids                 int
	AdresseDeRattachement string
	Label                 string `json:"Intitule"`
	StreetNumber          int    `json:"NumeroRue,string"`
	StreetName            string `json:"NomRue"`
	ComplementAdresse     string
	PostalCode            int    `json:"CodePostal,string"`
	City                  string `json:"Ville"`
	valElect              string
}

type MandatesJson struct {
	Mandate []MandateJson `json:"Mandat"`
}

type MandateJson struct {
	Uid             string
	ActeurId        string `json:"ActeurRef"`
	TermOffice      int    `json:"Legislature,string"`
	TypeOrgane      string
	StartDate       string          `json:"dateDebut"`
	PublicationDate string          `json:"datePublication"`
	EndDate         string          `json:"dateFin"`
	Precedence      int             `json:"preseance,string"`
	PrincipleNomin  int             `json:"nominPrincipale,string"`
	DataQuality     DataQualityJson `json:"infosQualite"`
	Body            BodyJson        `json:"organe"`
	Deputies        DeputiesJson    `json:"suppleants"`
	Election        ElectionJson
}

type DataQualityJson struct {
	QualityCode     string `json:"codeQualite"`
	QualityLabel    string `json:"libQualite"`
	QualityLabelSex string `json:"libQualiteSex"`
}

type BodyJson struct {
	RefBody string `json:"organeRef"`
}

type DeputiesJson struct {
	Deputy DeputyJson `json:"suppleant"`
}

type DeputyJson struct {
	StartDate string `json:"dateDebut"`
	EndDate   string `json:"dateFin"`
	RefDeputy string `json:"suppleantRef"`
}

type ElectionJson struct {
	MandateCause string    `json:"causeMandat"`
	DistrictRef  string    `json:"refCirconscription"`
	Place        PlaceJson `json:"lieu"`
}

type PlaceJson struct {
	Region        string
	TypeRegion    string `json:"regionType"`
	Department    string `json:"departement"`
	DepartmentNum int    `json:"numDepartement,string"`
	DistrictNum   int    `json:"numCirco,string"`
}
