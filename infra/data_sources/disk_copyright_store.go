package data_sources

import (
	"os"

	"golang.org/x/xerrors"
)

type DiskCopyrightStore struct {
}

func NewDiskCopyrightStore() *DiskCopyrightStore {
	return &DiskCopyrightStore{}
}

// wire

var singletonDiskCopyrightStore *DiskCopyrightStore = initSingletonDiskCopyrightStore()

func GetSingletonDiskCopyrightStore() *DiskCopyrightStore {
	return singletonDiskCopyrightStore
}

func initSingletonDiskCopyrightStore() *DiskCopyrightStore {
	return NewDiskCopyrightStore()
}

// methods
func (p *DiskCopyrightStore) CopyrightStorePath() string {
	return "copyrightstore"
}

func (p *DiskCopyrightStore) UnpackagedPath() string {
	return p.CopyrightStorePath() + "/unpackaged"
}

func (p *DiskCopyrightStore) PackagedPath() string {
	return p.CopyrightStorePath() + "/packaged"
}

func (p *DiskCopyrightStore) EnsureDirectory(path string) error {
	var err = os.MkdirAll(path, 0755)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	return nil
}

func (p *DiskCopyrightStore) EnsureCopyrightStoreDirectory() error {
	return p.EnsureDirectory(p.CopyrightStorePath())
}

func (p *DiskCopyrightStore) EnsureUnpackagedDirectory() error {
	return p.EnsureDirectory(p.UnpackagedPath())
}

func (p *DiskCopyrightStore) EnsurePackagedDirectory() error {
	return p.EnsureDirectory(p.PackagedPath())
}
