package repos

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/crclz/RightChain.cc/domain/domain_models"
	"github.com/crclz/RightChain.cc/infra/data_sources"
	"golang.org/x/xerrors"
)

type PackagedTreeRepository struct {
	diskCopyrightStore *data_sources.DiskCopyrightStore
}

func NewPackagedTreeRepository(
	diskCopyrightStore *data_sources.DiskCopyrightStore,
) *PackagedTreeRepository {
	return &PackagedTreeRepository{
		diskCopyrightStore: diskCopyrightStore,
	}
}

// wire

var singletonPackagedTreeRepository *PackagedTreeRepository = initSingletonPackagedTreeRepository()

func GetSingletonPackagedTreeRepository() *PackagedTreeRepository {
	return singletonPackagedTreeRepository
}

func initSingletonPackagedTreeRepository() *PackagedTreeRepository {
	return NewPackagedTreeRepository(data_sources.GetSingletonDiskCopyrightStore())
}

// methods

func (p *PackagedTreeRepository) GetPersistencePath(previousCommit string) string {
	var filePath = fmt.Sprintf("%v/%v.json", p.diskCopyrightStore.PackagedPath(), previousCommit)
	return filePath
}

func (p *PackagedTreeRepository) SavePackagedTree(ctx context.Context, tree *domain_models.PackagedTree) error {
	var err = p.diskCopyrightStore.EnsurePackagedDirectory()
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

func (p *PackagedTreeRepository) GetPackagedTreeByPreviousCommit(
	ctx context.Context, previousCommit string,
) (*domain_models.PackagedTree, error) {
	fileContent, err := ioutil.ReadFile(p.GetPersistencePath(previousCommit))
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	var result = &domain_models.PackagedTree{}
	err = json.Unmarshal(fileContent, result)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return result, nil
}

func (p *PackagedTreeRepository) GetAllPackagedTrees(ctx context.Context) ([]*domain_models.PackagedTree, error) {

	files, err := filepath.Glob(p.diskCopyrightStore.PackagedPath() + "/*.json")
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	var results []*domain_models.PackagedTree

	for _, filename := range files {
		fileContent, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}

		var result = &domain_models.PackagedTree{}
		err = json.Unmarshal(fileContent, result)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}

		results = append(results, result)
	}

	return results, nil
}
