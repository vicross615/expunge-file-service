package service

import (
	"fmt"
	"github.com/DkreativeCoders/expunge-file-service/pkg/domain"
	"log"
	"os"
	"path/filepath"
)

type ProcessFileUsingRecursiveDepth struct {
}

func NewProcessFileUsingRecursiveDepth() *ProcessFileUsingRecursiveDepth {
	return &ProcessFileUsingRecursiveDepth{}
}

func (p ProcessFileUsingRecursiveDepth) prepareFile(generalConfig domain.GeneralConfig,
	serviceConfig domain.ServiceConfig,
	fileProcessState *domain.FileProcessState) {

	if serviceConfig.EnableRecursiveDepth{
		//perform recursive check
		setOfExcludedFolder := p.getExcludedFolderAsSet(serviceConfig.ExcludeFolders)

		err := filepath.Walk(serviceConfig.RootPath,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if info.IsDir() && setOfExcludedFolder[info.Name()]{
					fmt.Print("found and skipping excludedFolder==>",info.Name())
					return filepath.SkipDir
				}

				if !info.IsDir(){
					fileProcessState.SetOfFilesPath[path]=true
				}
				//fmt.Println(path, "====>", info.Size(), "date===>", info.ModTime())
				return nil
			})
		if err != nil {
			log.Println(err)
		}

		fmt.Println("After ProcessFileUsingRecursiveDepth")

		//fileProcessState.SetOfFilesPath=setOfFilePath
	}
}
func (p ProcessFileUsingRecursiveDepth) getExcludedFolderAsSet(excludeFolders []string) map[string]bool{
	excludedFolderAsSet := make(map[string]bool)
	for _, s := range excludeFolders {
		excludedFolderAsSet[s] = true
	}
	return excludedFolderAsSet
}




