package application

import (
	"context"
	"log"

	"github.com/crclz/RightChain.cc/domain/domain_models"
	"github.com/crclz/RightChain.cc/domain/domain_services"
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

	for _, unpackagedTree := range unpackagedTrees {
		recordResponse, err := p.rightchainCenterService.OutOfBoxGetRecord(ctx, unpackagedTree.RecordFetchToken)
		if err != nil {
			return xerrors.Errorf(": %w", err)
		}

		if recordResponse.TransactionId == "" {
			log.Printf("still not packaged: %v", unpackagedTree.PreviousCommit)
			continue
		}

		// do package
		var slimTree = recordResponse.SlimTree
		target, targetParent := slimTree.FindNode(func(x *domain_models.RecipeNode) bool {
			return x.Literal == unpackagedTree.PartialTree.GetOutput()
		})

		if target == nil {
			return xerrors.Errorf("Cannot match literal when processing: %v", unpackagedTree.PreviousCommit)
		}

		if (targetParent.Left == target) != (targetParent.Right == target) {
			panic("find node logic error")
		}

		if targetParent.Left == target {
			targetParent.Left = unpackagedTree.PartialTree
		} else {
			targetParent.Right = unpackagedTree.PartialTree
		}

		slimTree.ClearCache()

		if recordResponse.RootOutput != slimTree.GetOutput() {
			return xerrors.Errorf("recordResponse output not mactch slimTree output: %v", unpackagedTree.PreviousCommit)
		}

		// success. TODO: save packagedTree ,
		// TODO: 删除. (暂时不删除)
		var packagedTree = &domain_models.PackagedTree{
			PreviousCommit: unpackagedTree.PreviousCommit,
			TransactionId:  recordResponse.TransactionId,
			RootOutput:     recordResponse.RootOutput,
			Tree:           slimTree,
		}

		log.Printf("TODO: save pacakgedTree. %v, %v, %v", packagedTree.PreviousCommit, packagedTree.RootOutput, packagedTree.Tree.GetOutput())
	}

	return nil
}
