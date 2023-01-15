package repos

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/crclz/RightChain.cc/domain/domain_models"
	"github.com/crclz/RightChain.cc/infra/data_sources"
	"golang.org/x/xerrors"
)

type UnpackagedTreeRepository struct {
	diskCopyrightStore *data_sources.DiskCopyrightStore
}

func NewUnpackagedTreeRepository(
	diskCopyrightStore *data_sources.DiskCopyrightStore,
) *UnpackagedTreeRepository {
	return &UnpackagedTreeRepository{
		diskCopyrightStore: diskCopyrightStore,
	}
}

// wire

var singletonUnpackagedTreeRepository *UnpackagedTreeRepository = initSingletonUnpackagedTreeRepository()

func GetSingletonUnpackagedTreeRepository() *UnpackagedTreeRepository {
	return singletonUnpackagedTreeRepository
}

func initSingletonUnpackagedTreeRepository() *UnpackagedTreeRepository {
	return NewUnpackagedTreeRepository(data_sources.GetSingletonDiskCopyrightStore())
}

// methods

func (p *UnpackagedTreeRepository) GetPersistencePath(previousCommit string) string {
	var filePath = fmt.Sprintf("%v/%v.json", p.diskCopyrightStore.UnpackagedPath(), previousCommit)
	return filePath
}

func (p *UnpackagedTreeRepository) SaveUnpackagedTree(ctx context.Context, tree *domain_models.UnpackagedTree) error {
	var err = p.diskCopyrightStore.EnsureUnpackagedDirectory()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	fileContent, err := json.MarshalIndent(tree, "", "\t")
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	err = ioutil.WriteFile(p.GetPersistencePath(tree.PreviousCommit), fileContent, 0644)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (p *UnpackagedTreeRepository) GetUnpackagedTreeByPreviousCommit(
	ctx context.Context, previousCommit string,
) (*domain_models.UnpackagedTree, error) {
	fileContent, err := ioutil.ReadFile(p.GetPersistencePath(previousCommit))
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	var result = &domain_models.UnpackagedTree{}
	err = json.Unmarshal(fileContent, result)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return result, nil
}

func (p *UnpackagedTreeRepository) GetAllUnpackagedTrees(ctx context.Context) ([]*domain_models.UnpackagedTree, error) {

	files, err := filepath.Glob(p.diskCopyrightStore.UnpackagedPath() + "/*.json")
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	var results []*domain_models.UnpackagedTree

	for _, filename := range files {
		fileContent, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}

		var result = &domain_models.UnpackagedTree{}
		err = json.Unmarshal(fileContent, result)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}

		results = append(results, result)
	}

	return results, nil
}

func (p *UnpackagedTreeRepository) Remove(
	ctx context.Context, tree *domain_models.UnpackagedTree,
) error {
	var filename = p.GetPersistencePath(tree.PreviousCommit)
	var err = os.Remove(filename)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
