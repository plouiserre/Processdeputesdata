-- créer la base de données 
CREATE DATABASE IF NOT EXISTS PROCESSDEPUTES;

-- créer l'utilisateur pour la BDD
DROP USER 'ProcessDeputesData'@'localhost' ;
CREATE USER 'ProcessDeputesData'@'localhost' IDENTIFIED BY 'ASimpleP@ssW0rd' ;

-- ajout de droits
GRANT DELETE ON PROCESSDEPUTES.* TO 'ProcessDeputesData'@'localhost' ;
GRANT INSERT ON PROCESSDEPUTES.* TO 'ProcessDeputesData'@'localhost' ;
GRANT UPDATE ON PROCESSDEPUTES.* TO 'ProcessDeputesData'@'localhost' ;
GRANT SELECT ON PROCESSDEPUTES.* TO 'ProcessDeputesData'@'localhost' ;

USE PROCESSDEPUTES;

-- créer les tables
-- 1 Election
CREATE TABLE IF NOT EXISTS  Election(
    ElectionId INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    MandateCause VARCHAR(50) NOT NULL,
    Region VARCHAR (50) NOT NULL,
    TypeRegion VARCHAR (50) NOT NULL,
    Department VARCHAR (50) NOT NULL,
    DepartmentNum INT NOT NULL,
    DistrictNum INT NOT NULL
);

-- 2 Deputy
CREATE TABLE IF NOT EXISTS  Deputy(
    DeputyId INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    StartDate DATETIME NOT NULL,
    EndDate DATETIME NOT NULL,
    RefDeputy VARCHAR (50) NOT NULL UNIQUE
);

-- 3 Congressman
CREATE TABLE IF NOT EXISTS  Congressman(
    CongressManId INT PRIMARY KEY NOT NULL  AUTO_INCREMENT,
    CongressManUid VARCHAR(50) NOT NULL UNIQUE,
    Civility VARCHAR(3) NOT NULL,
    FirstName VARCHAR(50) NOT NULL,
    LastName VARCHAR (50) NOT NULL,
    Alpha VARCHAR (50) NOT NULL,
    Trigramme VARCHAR (50) NOT NULL,
    BirthDate DATETIME NOT NULL,
    BirthCity VARCHAR (50) NOT NULL,
    BirthDepartment VARCHAR (50) NOT NULL,
    BirthCountry VARCHAR (50) NOT NULL,
    JobTitle VARCHAR (50) NOT NULL,
    CatSocPro VARCHAR (50) NOT NULL,
    FamSocPro VARCHAR (50) NOT NULL
);

-- 4 Mandate
CREATE TABLE IF NOT EXISTS  Mandate(
    MandateId INT PRIMARY KEY NOT NULL  AUTO_INCREMENT,
    MandateUid VARCHAR(50) NOT NULL UNIQUE,
    TermOffice INT NOT NULL,
    TypeOrgane VARCHAR(50) NOT NULL,
    StartDate DATETIME NOT NULL,
    EndDate DATETIME,
    Precedence INT NOT NULL,
    PrincipleNoming INT NOT NULL,
    QualityCode VARCHAR(50),
    QualityLabel VARCHAR(50),
    QualityLabelSex VARCHAR(50),
    RefBody VARCHAR(50) NOT NULL,
    DeputyId INT,
    ElectionId INT,
    CongressManId INT NOT NULL,
    FOREIGN KEY (DeputyId) REFERENCES Deputy(DeputyId),
    FOREIGN KEY (ElectionId) REFERENCES Election(ElectionId),
    FOREIGN KEY (CongressManId) REFERENCES CongressMan(CongressManId)
);