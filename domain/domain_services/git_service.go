package domain_services

import (
	"context"
	"log"
	"os/exec"

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
		return xerrors.Errorf(": %w", err)
	}

	log.Printf("git version output: %v", string(result))

	return nil
}
