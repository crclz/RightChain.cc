package application_test

import (
	"context"
	"os"
	"testing"

	"github.com/crclz/rightchain.cc/application"
	"github.com/crclz/rightchain.cc/domain/utils"
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

func TestDefaultController_FetchAllUnpackagedIndexs_happy_1(t *testing.T) {
	var assert = utils.AnyAssert(t)

	currentDir, err := os.Getwd()
	assert.NoError(err)
	defer os.Chdir(currentDir)
	os.Chdir("..")

	// arrange
	var ctx = context.TODO()
	var controller = application.GetSingletonDefaultController()

	// act
	err = controller.FetchAllUnpackagedIndexs(ctx)

	// assert
	assert.NoError(err)
}

func TestDefaultController_GenerateProof_happy_1(t *testing.T) {
	var assert = utils.AnyAssert(t)

	currentDir, err := os.Getwd()
	assert.NoError(err)
	defer os.Chdir(currentDir)
	os.Chdir("..")

	// arrange
	var ctx = context.TODO()
	var controller = application.GetSingletonDefaultController()

	// act
	err = controller.GenerateProof(ctx, []string{"readme.md", "setup.py"}, false)

	// assert
	assert.NoError(err)
}
