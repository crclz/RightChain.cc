package domain_services

import (
	"context"

	"github.com/crclz/rightchain.cc/domain/dtos"
	"github.com/imroc/req/v3"
	"golang.org/x/xerrors"
)

type RightchainCenterService struct {
}

func NewRightchainCenterService() *RightchainCenterService {
	return &RightchainCenterService{}
}

// wire

var singletonRightchainCenterService *RightchainCenterService = initSingletonRightchainCenterService()

func GetSingletonRightchainCenterService() *RightchainCenterService {
	return singletonRightchainCenterService
}

func initSingletonRightchainCenterService() *RightchainCenterService {
	return NewRightchainCenterService()
}

// methods

func (p *RightchainCenterService) BaseUrl() string {
	const useLocal = false
	if useLocal {
		return "http://localhost:5071"
	} else {
		return "https://rightchain.cc"
	}
}

func (p *RightchainCenterService) OutOfBoxCreateRecord(ctx context.Context, content string) (*dtos.OutOfBoxCreateRecordResponse, error) {
	var url = p.BaseUrl() + "/api/out-of-box/create-record"
	httpResponse, err := req.R().
		SetBodyJsonMarshal(map[string]interface{}{"content": content}).
		SetContext(ctx).
		Post(url)

	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	var response = &dtos.OutOfBoxCreateRecordResponse{}
	err = httpResponse.UnmarshalJson(response)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return response, nil
}

func (p *RightchainCenterService) OutOfBoxGetRecord(ctx context.Context, token string) (*dtos.OutOfBoxGetRecordResponse, error) {
	var url = p.BaseUrl() + "/api/out-of-box/get-record"
	httpResponse, err := req.R().
		SetQueryParam("token", token).
		SetContext(ctx).
		Get(url)

	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	var response = &dtos.OutOfBoxGetRecordResponse{}
	err = httpResponse.UnmarshalJson(response)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return response, nil
}
