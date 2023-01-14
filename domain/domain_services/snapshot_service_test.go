package domain_services_test

import (
	"context"
	"os"
	"testing"

	"github.com/crclz/RightChain.cc/domain/domain_services"
	"github.com/crclz/RightChain.cc/domain/utils"
)

func TestSnapshotService_ListFiles_happy_case_1(t *testing.T) {
	// arrange
	var assert = utils.AnyAssert(t)
	var ctx = context.TODO()
	var snaphotService = domain_services.GetSingletonSnaphotService()

	currentDir, err := os.Getwd()
	assert.NoError(err)
	defer os.Chdir(currentDir)
	os.Chdir("../..")

	// act
	result, err := snaphotService.ListFiles(ctx)
	assert.NoError(err)

	// assert
	t.Logf("Result things: %v", result)
}
