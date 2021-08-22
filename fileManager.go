package main

import(
	"log"
	"os"
	"io/ioutil"
)

type fileManager struct{
	nameFile string
	contentFile string
}

func (fileManager *fileManager) getContentFile(){
	file, err := os.OpenFile(fileManager.nameFile,  os.O_RDONLY,0666)
	if err != nil{
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil{
		log.Fatal(err)
	}
	displayContnet := string(content)
	fileManager.contentFile = displayContnet
}