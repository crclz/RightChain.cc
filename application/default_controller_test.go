package application_test

import (
	"context"
	"os"
	"testing"

	"github.com/crclz/RightChain.cc/application"
	"github.com/crclz/RightChain.cc/domain/utils"
)

func TestDefaultController_TakeSnapshotAndUpload_happy_1(t *testing.T) {
	var assert = utils.AnyAssert(t)

	currentDir, err := os.Getwd()
	assert.NoError(err)
	defer os.Chdir(currentDir)
	os.Chdir("..")

	// arrange
	var ctx = context.TODO()
	var controller = application.GetSingletonDefaultController()

	// act
	err = controller.TakeSnapshotAndUpload(ctx)

	// assert
	assert.NoError(err)
}
