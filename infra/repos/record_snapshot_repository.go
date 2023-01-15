package repos

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/crclz/RightChain.cc/domain/domain_models"
	"github.com/crclz/RightChain.cc/infra/data_sources"
	"golang.org/x/xerrors"
)

type RecordSnapshotRepository struct {
	diskCopyrightStore *data_sources.DiskCopyrightStore
}

func NewRecordSnapshotRepository(
	diskCopyrightStore *data_sources.DiskCopyrightStore,
) *RecordSnapshotRepository {
	return &RecordSnapshotRepository{
		diskCopyrightStore: diskCopyrightStore,
	}
}

// wire

var singletonRecordSnapshotRepository *RecordSnapshotRepository = initSingletonRecordSnapshotRepository()

func GetSingletonRecordSnapshotRepository() *RecordSnapshotRepository {
	return singletonRecordSnapshotRepository
}

func initSingletonRecordSnapshotRepository() *RecordSnapshotRepository {
	return NewRecordSnapshotRepository(data_sources.GetSingletonDiskCopyrightStore())
}

// methods
func (p *RecordSnapshotRepository) SaveSnapshot(ctx context.Context, snapshot *domain_models.RepositorySnapshot) error {
	var err = p.diskCopyrightStore.EnsureCopyrightStoreDirectory()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	snapshotJson, err := json.MarshalIndent(snapshot, "", "\t")
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	err = ioutil.WriteFile(p.diskCopyrightStore.CopyrightStorePath()+"/snapshot.json", snapshotJson, 0644)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
