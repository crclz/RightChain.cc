package domain_services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"

	"github.com/crclz/RightChain.cc/domain/domain_models"
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
			// log.Printf("Ignore pattern match: %v", filename)
			continue
		}

		newResults = append(newResults, filename)
	}

	return newResults, nil
}

func (p *SnaphotService) TakeSnapshot(ctx context.Context) (*domain_models.RepositorySnapshot, error) {
	commitHash, err := p.gitService.GetPreviousCommitHash(ctx)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	filenames, err := p.ListFiles(ctx)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	var fileSnapshots []*domain_models.FileSnapshot

	for _, filename := range filenames {
		hashString, err := p.Sha256(ctx, filename)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}

		fileSnapshots = append(fileSnapshots, &domain_models.FileSnapshot{
			Filename: filename,
			Hash:     hashString,
		})
	}

	return &domain_models.RepositorySnapshot{
		PreviousCommit: commitHash,
		FileSnapshots:  fileSnapshots,
	}, nil
}

func (p *SnaphotService) Sha256(ctx context.Context, filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", xerrors.Errorf(": %w", err)
	}
	defer f.Close()

	h := sha256.New()
	_, err = io.Copy(h, f)

	if err != nil {
		return "", xerrors.Errorf(": %w", err)
	}

	var hashBytes = h.Sum(nil)

	return hex.EncodeToString(hashBytes), nil
}
