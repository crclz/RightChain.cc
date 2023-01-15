package repos

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

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
func (p *UnpackagedTreeRepository) SaveUnpackagedTree(ctx context.Context, tree *domain_models.UnpackagedTree) error {
	var err = p.diskCopyrightStore.EnsureCopyrightStoreDirectory()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	fileContent, err := json.MarshalIndent(tree, "", "\t")
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	var filePath = fmt.Sprintf("%v/%v.json", p.diskCopyrightStore.UnpackagedPath(), tree.PreviousCommit)

	err = ioutil.WriteFile(filePath, fileContent, 0644)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
