package main

import(
	"os"
	"io/ioutil"
)

type fileManager struct{
	nameFile string
	contentFile string
	LogManager LogManager
}

func (fileManager *fileManager) getContentFile(){
	file, err := os.OpenFile(fileManager.nameFile,  os.O_RDONLY,0666)
	if err != nil{
		fileManager.LogManager.WriteErrorLog(err.Error())
	}
	content, err := ioutil.ReadAll(file)
	if err != nil{
		fileManager.LogManager.WriteErrorLog(err.Error())
	}
	displayContnet := string(content)
	fileManager.contentFile = displayContnet
}