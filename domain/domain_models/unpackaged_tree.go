package domain_models

type UnpackagedTree struct {
	PreviousCommit   string      `json:"previousCommit"`
	RecordFetchToken string      `json:"recordFetchToken"`
	PartialTree      *RecipeNode `json:"partialTree"`
}
