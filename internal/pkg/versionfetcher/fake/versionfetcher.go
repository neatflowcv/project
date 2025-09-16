package fakeversionfetcher

import (
	"context"

	versionfetcher "github.com/neatflowcv/project/internal/pkg/versionfetcher/core"
)

var _ versionfetcher.VersionFetcher = (*FakeVersionFetcher)(nil)

type FakeVersionFetcher struct {
	GoVersion           string
	GolangciLintVersion string
}

func NewFakeVersionFetcher() *FakeVersionFetcher {
	return &FakeVersionFetcher{
		GoVersion:           "1.25.1",
		GolangciLintVersion: "2.4.0",
	}
}

func (f *FakeVersionFetcher) FetchGoVersion(ctx context.Context) (string, error) {
	return f.GoVersion, nil
}

func (f *FakeVersionFetcher) FetchGolangciLintVersion(ctx context.Context) (string, error) {
	return f.GolangciLintVersion, nil
}
