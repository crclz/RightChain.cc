package domain_services

import "github.com/crclz/RightChain.cc/domain/domain_models"

type TreeService struct {
	gitService *GitService
}

func NewTreeService(
	gitService *GitService,
) *TreeService {
	return &TreeService{
		gitService: gitService,
	}
}

// wire

var singletonTreeService *TreeService = initSingletonTreeService()

func GetSingletonTreeService() *TreeService {
	return singletonTreeService
}

func initSingletonTreeService() *TreeService {
	return NewTreeService(
		GetSingletonGitService(),
	)
}

// methods
func (p *TreeService) BuildTreeFromSnapshot(snapshot *domain_models.FileSnapshot)
