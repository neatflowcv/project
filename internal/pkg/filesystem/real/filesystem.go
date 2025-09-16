package realfilesystem

import (
	"fmt"
	"os"

	filesystem "github.com/neatflowcv/project/internal/pkg/filesystem/core"
)

var _ filesystem.Filesystem = (*RealFilesystem)(nil)

type RealFilesystem struct {
}

func NewRealFilesystem() *RealFilesystem {
	return &RealFilesystem{}
}

func (f *RealFilesystem) CreateDirectory(path string) error {
	err := os.Mkdir(path, 0750) //nolint:mnd
	if err != nil {
		return fmt.Errorf("create directory %s: %w", path, err)
	}

	return nil
}

func (f *RealFilesystem) CreateFile(path string, content []byte) error {
	err := os.WriteFile(path, content, 0600) //nolint:mnd
	if err != nil {
		return fmt.Errorf("create file %s: %w", path, err)
	}

	return nil
}
