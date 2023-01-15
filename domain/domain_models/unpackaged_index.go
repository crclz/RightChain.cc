package domain_models

type UnpackagedIndex struct {
	PreviousCommit   string      `json:"previousCommit"`
	RecordFetchToken string      `json:"recordFetchToken"`
	PartialTree      *RecipeNode `json:"partialTree"`
}
