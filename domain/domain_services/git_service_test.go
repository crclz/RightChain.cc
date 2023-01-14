package domain_services_test

import (
	"context"
	"testing"

	"github.com/crclz/RightChain.cc/domain/domain_services"
	"github.com/crclz/RightChain.cc/domain/utils"
)

func TestGitService_CheckGitInstallation_return_nil_when_git_installed(t *testing.T) {
	// arrange
	var assert = utils.AnyAssert(t)
	var ctx = context.TODO()
	var gitService = domain_services.GetSingletonGitService()

	// act
	var err = gitService.CheckGitInstallation(ctx)

	// assert
	assert.NoError(err)
}
