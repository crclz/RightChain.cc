package repos

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/crclz/rightchain.cc/domain/domain_models"
	"github.com/crclz/rightchain.cc/infra/data_sources"
	"golang.org/x/xerrors"
)

type PackagedIndexRepository struct {
	diskCopyrightStore *data_sources.DiskCopyrightStore
}

func NewPackagedIndexRepository(
	diskCopyrightStore *data_sources.DiskCopyrightStore,
) *PackagedIndexRepository {
	return &PackagedIndexRepository{
		diskCopyrightStore: diskCopyrightStore,
	}
}

// wire

var singletonPackagedIndexRepository *PackagedIndexRepository = initSingletonPackagedIndexRepository()

func GetSingletonPackagedIndexRepository() *PackagedIndexRepository {
	return singletonPackagedIndexRepository
}

func initSingletonPackagedIndexRepository() *PackagedIndexRepository {
	return NewPackagedIndexRepository(data_sources.GetSingletonDiskCopyrightStore())
}

// methods

func (p *PackagedIndexRepository) GetPersistencePath(previousCommit string) string {
	var filePath = fmt.Sprintf("%v/%v.json", p.diskCopyrightStore.PackagedPath(), previousCommit)
	return filePath
}

func (p *PackagedIndexRepository) SavePackagedIndex(ctx context.Context, tree *domain_models.PackagedIndex) error {
	var err = p.diskCopyrightStore.EnsurePackagedDirectory()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	fileContent, err := json.MarshalIndent(tree, "", "\t")
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	err = os.WriteFile(p.GetPersistencePath(tree.PreviousCommit), fileContent, 0644)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (p *PackagedIndexRepository) GetPackagedIndexByPreviousCommit(
	ctx context.Context, previousCommit string,
) (*domain_models.PackagedIndex, error) {
	fileContent, err := os.ReadFile(p.GetPersistencePath(previousCommit))

	if os.IsNotExist(err) {
		return nil, nil
	}

	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	var result = &domain_models.PackagedIndex{}
	err = json.Unmarshal(fileContent, result)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return result, nil
}

func (p *PackagedIndexRepository) GetAllPackagedIndexs(ctx context.Context) ([]*domain_models.PackagedIndex, error) {

	files, err := filepath.Glob(p.diskCopyrightStore.PackagedPath() + "/*.json")
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	var results []*domain_models.PackagedIndex

	for _, filename := range files {
		fileContent, err := os.ReadFile(filename)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}

		var result = &domain_models.PackagedIndex{}
		err = json.Unmarshal(fileContent, result)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}

		results = append(results, result)
	}

	return results, nil
}
