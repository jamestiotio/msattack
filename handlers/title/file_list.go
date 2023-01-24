package title

type FileData struct {
	FileName string `json:"file_name"`
	FileSize string `json:"file_size"`
	Hash     string `json:"hash"`
	Target   string `json:"target"`
	URL      string `json:"url"`
}

func generateFileList() []FileData {
	files := []FileData{}

	return files
}
