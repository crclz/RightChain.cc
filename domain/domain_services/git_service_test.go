package domain_services_test

import (
	"context"
	"testing"

	"github.com/bytedance/mockey"
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

func TestGitService_CheckGitInstallation_return_error_when_git_not_installed(t *testing.T) {
	// arrange
	var assert = utils.AnyAssert(t)

	mockey.PatchConvey("", t, func() {
		var ctx = context.TODO()
		var gitService = domain_services.GetSingletonGitService()

		mockey.Mock((*domain_services.GitService).GitCommandName).Return("git-mock-233").Build()

		// act
		var err = gitService.CheckGitInstallation(ctx)

		// assert
		assert.Error(err)
		assert.ErrorContains(err, "call git version error")
	})
}
