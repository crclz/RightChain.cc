package domain_services

import (
	"github.com/crclz/RightChain.cc/domain/domain_models"
	"github.com/crclz/RightChain.cc/domain/utils"
)

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

type RecipeNodeQueue []*domain_models.RecipeNode

func (p *RecipeNodeQueue) Count() int {
	return len(*p)
}

func (p *RecipeNodeQueue) Enqueue(x *domain_models.RecipeNode) {
	*p = append(*p, x)
}

func (p *RecipeNodeQueue) Dequeue() *domain_models.RecipeNode {
	if p.Count() == 0 {
		panic("queue is empty")
	}

	var x = (*p)[0]
	*p = (*p)[1:]

	return x
}

// methods
func (p *TreeService) BuildTreeFromSnapshot(snapshot *domain_models.RepositorySnapshot) *domain_models.RecipeNode {
	var nodeQueue RecipeNodeQueue

	for _, file := range snapshot.FileSnapshots {
		nodeQueue.Enqueue(&domain_models.RecipeNode{
			Literal: file.Hash,
		})
		nodeQueue.Enqueue(&domain_models.RecipeNode{
			Literal: utils.GenerateSalt(8),
		})
	}

	for nodeQueue.Count() > 1 {
		var tempQueue RecipeNodeQueue

		for nodeQueue.Count() > 0 {
			var x = nodeQueue.Dequeue()
			if nodeQueue.Count() == 0 {
				tempQueue.Enqueue(x)
			} else {
				var y = nodeQueue.Dequeue()
				tempQueue.Enqueue(&domain_models.RecipeNode{
					Left: x, Right: y,
				})
			}
		}

		nodeQueue = tempQueue
	}

	if nodeQueue.Count() != 1 {
		panic("logic error")
	}

	return nodeQueue.Dequeue()
}
