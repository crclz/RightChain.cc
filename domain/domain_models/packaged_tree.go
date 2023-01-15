package domain_models

type PackagedTree struct {
	PreviousCommit string      `json:"previousCommit"`
	TransactionId  string      `json:"transactionId"`
	RootOutput     string      `json:"rootOutput"`
	Tree           *RecipeNode `json:"tree"`
}
