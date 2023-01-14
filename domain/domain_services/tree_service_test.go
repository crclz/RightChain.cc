package domain_services_test

import (
	"testing"

	"github.com/crclz/RightChain.cc/domain/domain_models"
	"github.com/crclz/RightChain.cc/domain/domain_services"
	"github.com/crclz/RightChain.cc/domain/utils"
)

func TestTreeService_BuildTreeFromSnapshot_happy_case_1(t *testing.T) {
	var assert = utils.AnyAssert(t)

	// arrange
	var treeService = domain_services.GetSingletonTreeService()

	var snapshot = &domain_models.RepositorySnapshot{
		FileSnapshots: []*domain_models.FileSnapshot{
			{Hash: "a"},
			{Hash: "b"},
			{Hash: "c"},
			{Hash: "d"},
			{Hash: "e"},
			{Hash: "f"},
			{Hash: "g"},
			{Hash: "h"},
		},
	}

	// act
	var root = treeService.BuildTreeFromSnapshot(snapshot)

	// assert
	assert.NotNil(root)

	t.Logf("Root is: %v", utils.ToJson(root))
}
