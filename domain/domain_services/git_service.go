package domain_services

import (
	"context"
	"log"
	"os/exec"
	"strings"

	"golang.org/x/xerrors"
)

type GitService struct {
}

func NewGitService() *GitService {
	return &GitService{}
}

// wire

var singletonGitService *GitService = initSingletonGitService()

func GetSingletonGitService() *GitService {
	return singletonGitService
}

func initSingletonGitService() *GitService {
	return NewGitService()
}

// methods

func (p *GitService) GitCommandName() string {
	return "git"
}

func (p *GitService) CheckGitInstallation(ctx context.Context) error {
	result, err := exec.CommandContext(ctx, p.GitCommandName(), "version").Output()
	if err != nil {
		return xerrors.Errorf("call git version error. output: %v, error: %w", string(result), err)
	}

	log.Printf("git version output: %v", string(result))

	return nil
}

func (p *GitService) GetPreviousCommitHash(ctx context.Context) (string, error) {
	// git log --pretty=format:%H -1
	result, err := exec.CommandContext(ctx,
		p.GitCommandName(),
		"log", "--pretty=format:%H", "-1").Output()

	if err != nil {
		return "", xerrors.Errorf(
			"Get commit hash failed. Please make sure this repo has at lease 1 commit. output: %v, error: %w",
			string(result), err)
	}

	var commitHash = string(result)
	commitHash = strings.ToLower(commitHash)

	return commitHash, nil
}
