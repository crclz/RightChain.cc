package domain_models

type UnpackagedTree struct {
	RecordFetchToken string      `json:"recordFetchToken"`
	PartialTree      *RecipeNode `json:"partialTree"`
}
