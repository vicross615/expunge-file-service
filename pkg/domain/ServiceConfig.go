package domain

type ServiceConfig struct {
	ServiceName string
	RootPath string
	ExcludeFolders []string
	ExcludeExtensions []string
	ExcludeSpecificFileNames []string
	ExcludeFileNamesContaining []string
	UseGeneralConfig bool
	FileAgeLastModifiedInDays int
	EnableRecursiveDepth bool
	EnableFileMovementToBackupFolder bool
	PathToBackUpFolder string
}



