package domain_models

type Proof struct {
	TransactionId string            `json:"transactionId"`
	RootOutput    string            `json:"rootOutput"`
	FilenameMap   map[string]string `json:"filenameMap"`
	Tree          *RecipeNode       `json:"tree"`
}
