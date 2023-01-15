package domain_services_test

import (
	"context"
	"testing"

	"github.com/crclz/RightChain.cc/domain/domain_services"
	"github.com/crclz/RightChain.cc/domain/utils"
)

func TestRightchainCenterService_OutOfBoxCreateRecord_integrity_check(t *testing.T) {
	var assert = utils.AnyAssert(t)

	// arrange
	var ctx = context.TODO()
	var rightchainCenterService = domain_services.GetSingletonRightchainCenterService()
	var content = "aaaaaaa123123"

	// act
	response, err := rightchainCenterService.OutOfBoxCreateRecord(ctx, content)
	assert.NoError(err)

	// assert
	assert.NotZero(response.Token)
	assert.NotZero(response.RecordText)
	assert.NotNil(response.BridgeNode)

	t.Logf("Response is: %v", utils.ToJson(response))

	// integrity check
	assert.Equal(content, response.BridgeNode.Left.Literal)
	assert.Equal(response.RecordText, response.BridgeNode.GetOutput())
}

func TestRightchainCenterService_OutOfBoxGetRecord_happy_1(t *testing.T) {
	var assert = utils.AnyAssert(t)

	// arrange
	var ctx = context.TODO()
	var rightchainCenterService = domain_services.GetSingletonRightchainCenterService()
	var token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJyZWNvcmRJZCI6IjYzYzM4MjkyNjIxNDc4Y2IzYzVhOWY4MCJ9.HlPdPI0-qcp9JKja9NVXAfUCBSP7RcBoRPX4-EfHkmY"

	// act
	response, err := rightchainCenterService.OutOfBoxGetRecord(ctx, token)
	assert.NoError(err)

	// assert
	assert.NotZero(response.Id)
	assert.NotZero(response.Text)
	assert.NotZero(response.TransactionId)
	assert.NotZero(response.CreatedAt)
	assert.NotNil(response.SlimTree)
	assert.NotNil(response.RootOutput)

	t.Logf("Response is: %v", utils.ToJson(response))
}
