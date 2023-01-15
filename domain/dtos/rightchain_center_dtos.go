package dtos

import "github.com/crclz/RightChain.cc/domain/domain_models"

type OutOfBoxCreateRecordResponse struct {
	Token      string                    `json:"token"`
	BridgeNode *domain_models.RecipeNode `json:"bridgeNode"`
	RecordText string                    `json:"recordText"`
}
