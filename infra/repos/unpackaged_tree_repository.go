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

type UnpackagedIndexRepository struct {
	diskCopyrightStore *data_sources.DiskCopyrightStore
}

func NewUnpackagedIndexRepository(
	diskCopyrightStore *data_sources.DiskCopyrightStore,
) *UnpackagedIndexRepository {
	return &UnpackagedIndexRepository{
		diskCopyrightStore: diskCopyrightStore,
	}
}

// wire

var singletonUnpackagedIndexRepository *UnpackagedIndexRepository = initSingletonUnpackagedIndexRepository()

func GetSingletonUnpackagedIndexRepository() *UnpackagedIndexRepository {
	return singletonUnpackagedIndexRepository
}

func initSingletonUnpackagedIndexRepository() *UnpackagedIndexRepository {
	return NewUnpackagedIndexRepository(data_sources.GetSingletonDiskCopyrightStore())
}

// methods

func (p *UnpackagedIndexRepository) GetPersistencePath(previousCommit string) string {
	var filePath = fmt.Sprintf("%v/%v.json", p.diskCopyrightStore.UnpackagedPath(), previousCommit)
	return filePath
}

func (p *UnpackagedIndexRepository) SaveUnpackagedIndex(ctx context.Context, tree *domain_models.UnpackagedIndex) error {
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

func (p *UnpackagedIndexRepository) GetUnpackagedIndexByPreviousCommit(
	ctx context.Context, previousCommit string,
) (*domain_models.UnpackagedIndex, error) {
	fileContent, err := ioutil.ReadFile(p.GetPersistencePath(previousCommit))
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	var result = &domain_models.UnpackagedIndex{}
	err = json.Unmarshal(fileContent, result)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return result, nil
}

func (p *UnpackagedIndexRepository) GetAllUnpackagedIndexs(ctx context.Context) ([]*domain_models.UnpackagedIndex, error) {

	files, err := filepath.Glob(p.diskCopyrightStore.UnpackagedPath() + "/*.json")
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	var results []*domain_models.UnpackagedIndex

	for _, filename := range files {
		fileContent, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}

		var result = &domain_models.UnpackagedIndex{}
		err = json.Unmarshal(fileContent, result)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}

		results = append(results, result)
	}

	return results, nil
}

func (p *UnpackagedIndexRepository) Remove(
	ctx context.Context, tree *domain_models.UnpackagedIndex,
) error {
	var filename = p.GetPersistencePath(tree.PreviousCommit)
	var err = os.Remove(filename)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
