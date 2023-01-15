package application

import (
	"context"
	"log"

	"github.com/crclz/RightChain.cc/domain/domain_models"
	"github.com/crclz/RightChain.cc/domain/domain_services"
	"github.com/crclz/RightChain.cc/domain/utils"
	"github.com/crclz/RightChain.cc/infra/repos"
	"golang.org/x/xerrors"
)

type DefaultController struct {
	snaphotService           *domain_services.SnaphotService
	recordSnapshotRepository *repos.RecordSnapshotRepository
	rightchainCenterService  *domain_services.RightchainCenterService
	treeService              *domain_services.TreeService
	unpackagedTreeRepository *repos.UnpackagedTreeRepository
}

func NewDefaultController(
	snaphotService *domain_services.SnaphotService,
	recordSnapshotRepository *repos.RecordSnapshotRepository,
	rightchainCenterService *domain_services.RightchainCenterService,
	treeService *domain_services.TreeService,
	unpackagedTreeRepository *repos.UnpackagedTreeRepository,
) *DefaultController {
	return &DefaultController{
		snaphotService:           snaphotService,
		recordSnapshotRepository: recordSnapshotRepository,
		rightchainCenterService:  rightchainCenterService,
		treeService:              treeService,
		unpackagedTreeRepository: unpackagedTreeRepository,
	}
}

// wire

var singletonDefaultController *DefaultController = initSingletonDefaultController()

func GetSingletonDefaultController() *DefaultController {
	return singletonDefaultController
}

func initSingletonDefaultController() *DefaultController {
	return NewDefaultController(
		domain_services.GetSingletonSnaphotService(),
		repos.GetSingletonRecordSnapshotRepository(),
		domain_services.GetSingletonRightchainCenterService(),
		domain_services.GetSingletonTreeService(),
		repos.GetSingletonUnpackagedTreeRepository(),
	)
}

// methods

func (p *DefaultController) TakeSnapshotAndUpload(ctx context.Context) error {
	// take snapshot and save
	snapshot, err := p.snaphotService.TakeSnapshot(ctx)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	err = p.recordSnapshotRepository.SaveSnapshot(ctx, snapshot)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	// make tree from snapshot
	var partialTree = p.treeService.BuildTreeFromSnapshot(snapshot)

	// upload tree root node output
	createRecordResponse, err := p.rightchainCenterService.OutOfBoxCreateRecord(ctx, partialTree.GetOutput())
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	var bridgeNode = createRecordResponse.BridgeNode

	if bridgeNode.Left.Literal != partialTree.GetOutput() {
		panic("Server contract check fail")
	}

	bridgeNode.Left.Literal = ""
	bridgeNode.Left = partialTree

	// bridgeNode就是我们应当保存的
	var newPartialTree = bridgeNode
	newPartialTree.ClearCache()

	if bridgeNode.GetOutput() != createRecordResponse.RecordText {
		panic("Logic bug")
	}

	// make unpackaged tree and save
	var unpackagedTree = &domain_models.UnpackagedTree{
		PreviousCommit:   snapshot.PreviousCommit,
		RecordFetchToken: createRecordResponse.Token,
		PartialTree:      newPartialTree,
	}

	err = p.unpackagedTreeRepository.SaveUnpackagedTree(ctx, unpackagedTree)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (p *DefaultController) FetchAllUnpackagedTrees(ctx context.Context) error {
	unpackagedTrees, err := p.unpackagedTreeRepository.GetAllUnpackagedTrees(ctx)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	log.Printf("unpackagedTrees: %v", len(unpackagedTrees))

	panic(utils.ErrNotImplemented)
}
