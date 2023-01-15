package domain_models

type PackagedIndex struct {
	PreviousCommit string      `json:"previousCommit"`
	TransactionId  string      `json:"transactionId"`
	RootOutput     string      `json:"rootOutput"`
	Tree           *RecipeNode `json:"tree"`
}
