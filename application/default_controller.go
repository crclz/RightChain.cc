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
	snaphotService            *domain_services.SnaphotService
	recordSnapshotRepository  *repos.RecordSnapshotRepository
	rightchainCenterService   *domain_services.RightchainCenterService
	treeService               *domain_services.TreeService
	UnpackagedIndexRepository *repos.UnpackagedIndexRepository
	packagedIndexRepository   *repos.PackagedIndexRepository
	gitService                *domain_services.GitService
}

func NewDefaultController(
	snaphotService *domain_services.SnaphotService,
	recordSnapshotRepository *repos.RecordSnapshotRepository,
	rightchainCenterService *domain_services.RightchainCenterService,
	treeService *domain_services.TreeService,
	unpackagedIndexRepository *repos.UnpackagedIndexRepository,
	packagedIndexRepository *repos.PackagedIndexRepository,
	gitService *domain_services.GitService,
) *DefaultController {
	return &DefaultController{
		snaphotService:            snaphotService,
		recordSnapshotRepository:  recordSnapshotRepository,
		rightchainCenterService:   rightchainCenterService,
		treeService:               treeService,
		UnpackagedIndexRepository: unpackagedIndexRepository,
		packagedIndexRepository:   packagedIndexRepository,
		gitService:                gitService,
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
		repos.GetSingletonUnpackagedIndexRepository(),
		repos.GetSingletonPackagedIndexRepository(),
		domain_services.GetSingletonGitService(),
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
	var unpackagedIndex = &domain_models.UnpackagedIndex{
		PreviousCommit:   snapshot.PreviousCommit,
		RecordFetchToken: createRecordResponse.Token,
		PartialTree:      newPartialTree,
	}

	err = p.UnpackagedIndexRepository.SaveUnpackagedIndex(ctx, unpackagedIndex)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (p *DefaultController) FetchAllUnpackagedIndexs(ctx context.Context) error {
	UnpackagedIndexs, err := p.UnpackagedIndexRepository.GetAllUnpackagedIndexs(ctx)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	log.Printf("UnpackagedIndexs: %v", len(UnpackagedIndexs))

	for _, unpackagedIndex := range UnpackagedIndexs {
		recordResponse, err := p.rightchainCenterService.OutOfBoxGetRecord(ctx, unpackagedIndex.RecordFetchToken)
		if err != nil {
			return xerrors.Errorf(": %w", err)
		}

		if recordResponse.TransactionId == "" {
			log.Printf("still not packaged: %v", unpackagedIndex.PreviousCommit)
			continue
		}

		// do package
		var slimTree = recordResponse.SlimTree
		target, targetParent := slimTree.FindNode(func(x *domain_models.RecipeNode) bool {
			return x.Literal == unpackagedIndex.PartialTree.GetOutput()
		})

		if target == nil {
			return xerrors.Errorf("Cannot match literal when processing: %v", unpackagedIndex.PreviousCommit)
		}

		if (targetParent.Left == target) == (targetParent.Right == target) {
			panic("find node logic error")
		}

		if targetParent.Left == target {
			targetParent.Left = unpackagedIndex.PartialTree
		} else {
			targetParent.Right = unpackagedIndex.PartialTree
		}

		slimTree.ClearCache()

		if recordResponse.RootOutput != slimTree.GetOutput() {
			return xerrors.Errorf("recordResponse output not mactch slimTree output: %v", unpackagedIndex.PreviousCommit)
		}

		// success. TODO: save packagedIndex ,
		var packagedIndex = &domain_models.PackagedIndex{
			PreviousCommit: unpackagedIndex.PreviousCommit,
			TransactionId:  recordResponse.TransactionId,
			RootOutput:     recordResponse.RootOutput,
			Tree:           slimTree,
		}

		log.Printf("save pacakgedTree. %v, %v, %v", packagedIndex.PreviousCommit, packagedIndex.RootOutput, packagedIndex.Tree.GetOutput())
		err = p.packagedIndexRepository.SavePackagedIndex(ctx, packagedIndex)
		if err != nil {
			return xerrors.Errorf(": %w", err)
		}

		// TODO: 删除. (暂时不删除)
		err = p.UnpackagedIndexRepository.Remove(ctx, unpackagedIndex)
		if err != nil {
			return xerrors.Errorf(": %w", err)
		}
	}

	return nil
}

// proof: 由于package和commit的不同步问题，所以只能手动将hash拷贝
func (p *DefaultController) GenerateProof(ctx context.Context, files []string) error {
	// previousCommit, err := p.gitService.GetPreviousCommitHash(ctx)
	// if err != nil {
	// 	return xerrors.Errorf(": %w", err)
	// }

	panic(utils.ErrNotImplemented)

	// var packagedIndex = p.packagedIndexRepository.GetPackagedIndexByPreviousCommit(previousCommit)
}
