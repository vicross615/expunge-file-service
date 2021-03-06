package ProcessFile

import (
	"github.com/DkreativeCoders/expunge-file-service/pkg/domain"
	"github.com/kpango/glg"
	"log"
	"os"
	"path/filepath"
)

type Backup struct {
}

func NewProcessFileBackup() *Backup {
	return &Backup{}
}




func (p Backup) PrepareFile(generalConfig domain.GeneralConfig,
	serviceConfig domain.ServiceConfig,
	fileProcessState *domain.FileProcessState) {

	var enableFileMovementToBackupFolder bool
	var pathToBackupFile string

	if serviceConfig.UseGeneralConfig {
		enableFileMovementToBackupFolder=generalConfig.EnableFileMovementToBackupFolder
		pathToBackupFile = generalConfig.PathToBackupFile
	}else {
		enableFileMovementToBackupFolder=serviceConfig.EnableFileMovementToBackupFolder
		pathToBackupFile = serviceConfig.PathToBackUpFolder
	}


	if enableFileMovementToBackupFolder{
		p.performBackup(fileProcessState, pathToBackupFile)
		glg.Log("After ProcessFileBackup")
		glg.Log("SetOfFilesPathToBeDeleted ", fileProcessState.SetOfFilesPath)
	}


}

func (p Backup) performBackup(fileProcessState *domain.FileProcessState, backupFolderPath string) {

	if _, err := os.Stat(backupFolderPath); os.IsNotExist(err) {
		// path/to/whatever does not exist
		glg.Log("Backup Directory does not exist")
		glg.Log("Creating Backup directory")
		err = os.Mkdir(backupFolderPath, 0755)
		if err != nil {
			log.Fatal(err)
		}

	}

	for systemFilePath, _ := range fileProcessState.SetOfFilesPath {
		fileInfo, _ := os.Stat(systemFilePath)
		newPath := filepath.Join(backupFolderPath, fileInfo.Name())
		e := os.Rename(systemFilePath, newPath)
		if e != nil {
			log.Fatal(e)
		}
		
	}
}





