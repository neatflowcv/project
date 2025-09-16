package flow_test

import (
	"fmt"
	"testing"

	"github.com/neatflowcv/project/internal/app/flow"
	fakefilesystem "github.com/neatflowcv/project/internal/pkg/filesystem/fake"
	fakeversionfetcher "github.com/neatflowcv/project/internal/pkg/versionfetcher/fake"
	"github.com/stretchr/testify/require"
)

func TestNewProject(t *testing.T) {
	t.Parallel()

	const projectName = "test"

	filesystem := fakefilesystem.NewFakeFilesystem()
	fetcher := fakeversionfetcher.NewFakeVersionFetcher()
	service := flow.NewService(filesystem, fetcher)

	err := service.NewProject(t.Context(), "~", projectName, "test")

	require.NoError(t, err)
	require.Len(t, filesystem.Dirs, 8)
	require.True(t, filesystem.HasDirectory(fmt.Sprintf("~/workspace/%s", projectName))) //nolint:perfsprint
	require.True(t, filesystem.HasDirectory(fmt.Sprintf("~/workspace/%s/cmd", projectName)))
	require.True(t, filesystem.HasDirectory(fmt.Sprintf("~/workspace/%s/cmd/%s", projectName, projectName)))
	require.True(t, filesystem.HasFile(fmt.Sprintf("~/workspace/%s/cmd/%s/main.go", projectName, projectName)))
	require.True(t, filesystem.HasDirectory(fmt.Sprintf("~/workspace/%s/internal", projectName)))
	require.True(t, filesystem.HasDirectory(fmt.Sprintf("~/workspace/%s/internal/app", projectName)))
	require.True(t, filesystem.HasDirectory(fmt.Sprintf("~/workspace/%s/internal/app/flow", projectName)))
	require.True(t, filesystem.HasFile(fmt.Sprintf("~/workspace/%s/internal/app/flow/service.go", projectName)))
	require.True(t, filesystem.HasDirectory(fmt.Sprintf("~/workspace/%s/internal/pkg", projectName)))
	require.True(t, filesystem.HasDirectory(fmt.Sprintf("~/workspace/%s/internal/pkg/domain", projectName)))
	require.True(t, filesystem.HasFile(fmt.Sprintf("~/workspace/%s/.gitignore", projectName)))
	require.True(t, filesystem.HasFile(fmt.Sprintf("~/workspace/%s/.golangci.yaml", projectName)))
	require.True(t, filesystem.HasFile(fmt.Sprintf("~/workspace/%s/.pre-commit-config.yaml", projectName)))
	require.True(t, filesystem.HasFile(fmt.Sprintf("~/workspace/%s/go.mod", projectName)))
	require.True(t, filesystem.HasFile(fmt.Sprintf("~/workspace/%s/Makefile", projectName)))
}
