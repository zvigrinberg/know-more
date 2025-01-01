package core

import "os"

type ConfigManager struct {
	CommandsSourceFilePath string
}

func InitializeConfig() ConfigManager {
	config := ConfigManager{CommandsSourceFilePath: getSourceFileLocation()}
	return config
}

func getSourceFileLocation() string {
	filePath := os.Getenv("KNOW_MORE_FILE_PATH")
	// default source file path
	if filePath == "" {
		return "/tmp/trace"
	} else {
		return filePath
	}
}
