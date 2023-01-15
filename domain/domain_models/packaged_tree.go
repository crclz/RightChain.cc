package domain_models

type PackagedTree struct {
	PreviousCommit string      `json:"previousCommit"`
	TransactionId  string      `json:"transactionId"`
	RootHash       string      `json:"rootHash"`
	RepositoryTree *RecipeNode `json:"repositoryTree"`
}
