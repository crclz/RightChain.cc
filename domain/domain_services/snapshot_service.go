package domain_services

import (
	"context"
	"log"
	"path/filepath"

	"github.com/crclz/RightChain.cc/domain/domain_models"
	"github.com/crclz/RightChain.cc/domain/utils"
	"golang.org/x/xerrors"
)

type SnaphotService struct {
	gitService *GitService
}

func NewSnaphotService(
	gitService *GitService,
) *SnaphotService {
	return &SnaphotService{
		gitService: gitService,
	}
}

// wire

var singletonSnaphotService *SnaphotService = initSingletonSnaphotService()

func GetSingletonSnaphotService() *SnaphotService {
	return singletonSnaphotService
}

func initSingletonSnaphotService() *SnaphotService {
	return NewSnaphotService(
		GetSingletonGitService(),
	)
}

// methods

func (p *SnaphotService) TakeSnapshot(ctx context.Context) (*domain_models.RepositorySnapshot, error) {
	panic(utils.ErrNotImplemented)
}

func (p *SnaphotService) ListFiles(ctx context.Context) ([]string, error) {
	// include: git tracked and untracked files
	var result []string

	trackedFiles, err := p.gitService.ListTrackingFiles(ctx)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	result = append(result, trackedFiles...)

	untrackedFiles, err := p.gitService.ListUntrackedFiles(ctx)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	result = append(result, untrackedFiles...)

	// exclude: copyrightstore glob
	var ignorePattern = "copyrightstore/**"

	var newResults []string
	for _, filename := range result {
		matched, err := filepath.Match(ignorePattern, filename)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}

		if matched {
			log.Printf("Ignore pattern match: %v", filename)
			continue
		}

		newResults = append(newResults, filename)
	}

	return newResults, nil
}
