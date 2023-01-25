package managers

import (
	"os"

	"msattack/config"

	"github.com/bytedance/sonic"
	"github.com/rs/zerolog/log"
)

type FileData struct {
	FileName string `json:"file_name"`
	FileSize string `json:"file_size"`
	Hash     string `json:"hash"`
	// "target" is "0" for normal files and "1" for master table files
	Target string `json:"target"`
	URL    string `json:"url"`
}

func GenerateFileList() []FileData {
	// The backend server serves different sets of files depending on the current game version.
	// Currently, this function will only generate the latest file list for now.
	// Future work can be done to incorporate the mapping of game versions to file lists.

	configuration := config.GlobalConfig

	var files []FileData

	switch configuration.MasterVersion {
	case 7130000:
		fileListBytes, err := os.ReadFile(configuration.FileListFilename)
		if err != nil {
			log.Error().Err(err).Msg("Failed to read file list.")
			files = []FileData{}
		} else {
			err = sonic.Unmarshal(fileListBytes, &files)
			if err != nil {
				log.Error().Err(err).Msg("Failed to unmarshal file list.")
				files = []FileData{}
			}
		}
	default:
		files = []FileData{}
	}

	return files
}
