package domain_models

type RepositorySnapshot struct {
	PreviousCommit string          `json:"previousCommit"`
	FileSnapshots  []*FileSnapshot `json:"fileSnapshots"`
}

type FileSnapshot struct {
	Filename string `json:"filename"`
	Hash     string `json:"hash"`
}
