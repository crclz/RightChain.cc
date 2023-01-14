package domain_services

import (
	"context"

	"github.com/crclz/RightChain.cc/domain/domain_models"
	"github.com/crclz/RightChain.cc/domain/utils"
)

type SnaphotService struct {
}

func NewSnaphotService() *SnaphotService {
	return &SnaphotService{}
}

// wire

var singletonSnaphotService *SnaphotService = initSingletonSnaphotService()

func GetSingletonSnaphotService() *SnaphotService {
	return singletonSnaphotService
}

func initSingletonSnaphotService() *SnaphotService {
	return NewSnaphotService()
}

// methods

func (p *SnaphotService) TakeSnapshot(
	ctx context.Context, directory string,
) (*domain_models.RepositorySnapshot, error) {
	panic(utils.ErrNotImplemented)
}
