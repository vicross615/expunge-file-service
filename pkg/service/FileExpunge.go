package service

import (
	"fmt"
	"github.com/DkreativeCoders/expunge-file-service/pkg/domain"
	"github.com/DkreativeCoders/expunge-file-service/pkg/utils"
	"log"
)

type FileExpunge struct {
	fileJsonParser IFileJsonParse
	factoryProcessFile FactoryProcessFile
}

func NewFileExpunge(fileJsonParser IFileJsonParse, factoryProcessFile FactoryProcessFile) *FileExpunge {
	return &FileExpunge{fileJsonParser: fileJsonParser, factoryProcessFile: factoryProcessFile}
}




func (f FileExpunge) ExecuteDeleteTask() {
	pathToFileCleanerJson :=utils.GetPathToFileCleanerJson()
	fileCleanerJsonConfig,err:= f.fileJsonParser.ParseFileCleanerJson(pathToFileCleanerJson)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println("fileCleanerJsonConfig serviceConfigs", fileCleanerJsonConfig.ServiceConfigs)


	for _, serViceConfig := range fileCleanerJsonConfig.ServiceConfigs {

		fmt.Println(">>>> Starting Executing for ", serViceConfig.ServiceName)

		fileProcessState :=domain.NewFileProcessState()
		fileProcessState.SetOfFilesPath = make(map[string]bool)

		for _, processFile := range f.factoryProcessFile.ProcessFile{
			processFile.prepareFile(fileCleanerJsonConfig.GeneralConfig,serViceConfig,fileProcessState)
		}
		fileProcessState=nil
		fmt.Println(">>>> Done Executing for ", serViceConfig.ServiceName)
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

	}

	//serViceConfig:= fileCleanerJsonConfig.ServiceConfigs[0]

	//processFileNonRecursive := NewProcessFileNonRecursive()
	//processFileNonRecursive.prepareFile(fileCleanerJsonConfig.GeneralConfig,serViceConfig,fileProcessState)
	//
	//processFileUsingRecursiveDepth :=NewProcessFileUsingRecursiveDepth()
	//processFileUsingRecursiveDepth.prepareFile(fileCleanerJsonConfig.GeneralConfig,serViceConfig,fileProcessState)
	//
	//processFileRemoveExcludedExtension :=NewProcessFileExcludedExtension()
	//processFileRemoveExcludedExtension.prepareFile(fileCleanerJsonConfig.GeneralConfig,serViceConfig,fileProcessState)
	//
	//processFileRemoveSpecificFileName:=NewProcessFileExcludeSpecificFileName()
	//processFileRemoveSpecificFileName.prepareFile(fileCleanerJsonConfig.GeneralConfig,serViceConfig,fileProcessState)
	//
	//processFileExcludeFileOfConfigAge :=NewProcessFileExcludeFileOfConfigAge()
	//processFileExcludeFileOfConfigAge.prepareFile(fileCleanerJsonConfig.GeneralConfig,serViceConfig,fileProcessState)
	//
	//processFileBackup := NewProcessFileBackup()
	//processFileBackup.prepareFile(fileCleanerJsonConfig.GeneralConfig,serViceConfig,fileProcessState)
	//
	//processFileEradicate :=NewProcessFileEradicate()
	//processFileEradicate.prepareFile(fileCleanerJsonConfig.GeneralConfig,serViceConfig,fileProcessState)

	//fileProcessState=nil

	//fmt.Println("set of path",*fileProcessState)

	//-Walk through the whole file path and get all the file paths needed if
	//  ---"enableRecursiveDepth" : true,

	//	-here you exclude directories that are not meant to be searched
	//-remove excluded extensions
	//-remove excluded fileNames
	//-remove excludeFileNamesContaining
	//-check last modified date of the file ==  "fileAgeToBeDeleted": 120,
	//-check if "enableFileMovementToBackupFolder": true,
	//   --if true  get  "pathToBackUpFolder":"C:\\Users\\dell\\Documents\\daniel\\backup"
	//-move file to backupFolder
	//delete all the files


	//      "excludeExtensions": [".jar",".bat"],
	//      "excludeFileNames" : ["do-not-delete","delete-not"],
	//      "excludeFileNamesContaining" : ["tomcat","tomcat"],
	//panic("implement me")
}




