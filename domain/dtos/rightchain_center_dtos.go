package dtos

import "github.com/crclz/RightChain.cc/domain/domain_models"

type OutOfBoxCreateRecordResponse struct {
	Token      string                    `json:"token"`
	BridgeNode *domain_models.RecipeNode `json:"bridgeNode"`
	RecordText string                    `json:"recordText"`
}

type OutOfBoxGetRecordResponse struct {
	Id            string                    `json:"id"`
	Text          string                    `json:"text"`
	TransactionId string                    `json:"transactionId"`
	CreatedAt     int64                     `json:"createdAt"`
	SlimTree      *domain_models.RecipeNode `json:"slimTree"`
	RootOutput    string                    `json:"rootOutput"`
}
