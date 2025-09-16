package fakefilesystem

import filesystem "github.com/neatflowcv/project/internal/pkg/filesystem/core"

var _ filesystem.Filesystem = (*FakeFilesystem)(nil)

type FakeFilesystem struct {
	Dirs  map[string]any
	Files map[string][]byte
}

func NewFakeFilesystem() *FakeFilesystem {
	return &FakeFilesystem{
		Dirs:  make(map[string]any),
		Files: make(map[string][]byte),
	}
}

func (f *FakeFilesystem) CreateDirectory(path string) error {
	f.Dirs[path] = nil

	return nil
}

func (f *FakeFilesystem) HasDirectory(s string) bool {
	_, ok := f.Dirs[s]

	return ok
}

func (f *FakeFilesystem) CreateFile(path string, content []byte) error {
	f.Files[path] = content

	return nil
}

func (f *FakeFilesystem) HasFile(s string) bool {
	_, ok := f.Files[s]

	return ok
}
