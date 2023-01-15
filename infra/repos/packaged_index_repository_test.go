package repos_test

import (
	"context"
	"testing"

	"github.com/crclz/RightChain.cc/domain/utils"
	"github.com/crclz/RightChain.cc/infra/repos"
)

func TestPackagedIndexRepository_GetPackagedIndexByPreviousCommit_return_nil_when_not_exist(t *testing.T) {
	var assert = utils.AnyAssert(t)

	// arrange
	var ctx = context.TODO()

	var packagedIndexRepository = repos.GetSingletonPackagedIndexRepository()

	// act
	result, err := packagedIndexRepository.GetPackagedIndexByPreviousCommit(ctx, "asd")
	assert.NoError(err)

	// assert
	assert.Nil(result)
}
