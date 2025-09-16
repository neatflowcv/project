package flow

import (
	"fmt"

	filesystem "github.com/neatflowcv/project/internal/pkg/filesystem/core"
	"github.com/neatflowcv/project/internal/pkg/templates"
	versionfetcher "github.com/neatflowcv/project/internal/pkg/versionfetcher/core"
)

type Service struct {
	filesystem filesystem.Filesystem
	fetcher    versionfetcher.VersionFetcher
}

func NewService(filesystem filesystem.Filesystem, fetcher versionfetcher.VersionFetcher) *Service {
	return &Service{
		filesystem: filesystem,
		fetcher:    fetcher,
	}
}

func (s *Service) NewProject(userHome string, projectName string, moduleName string) error { //nolint:funlen
	dirs := []string{
		fmt.Sprintf("%s/workspace/%s", userHome, projectName),
		fmt.Sprintf("%s/workspace/%s/cmd", userHome, projectName),
		fmt.Sprintf("%s/workspace/%s/cmd/%s", userHome, projectName, projectName),
		fmt.Sprintf("%s/workspace/%s/internal", userHome, projectName),
		fmt.Sprintf("%s/workspace/%s/internal/app", userHome, projectName),
		fmt.Sprintf("%s/workspace/%s/internal/app/flow", userHome, projectName),
		fmt.Sprintf("%s/workspace/%s/internal/pkg", userHome, projectName),
		fmt.Sprintf("%s/workspace/%s/internal/pkg/domain", userHome, projectName),
	}

	for _, dir := range dirs {
		err := s.filesystem.CreateDirectory(dir)
		if err != nil {
			return fmt.Errorf("create directory %s: %w", dir, err)
		}
	}

	goVersion, err := s.fetcher.FetchGoVersion()
	if err != nil {
		return fmt.Errorf("fetch go version: %w", err)
	}

	golangciLintVersion, err := s.fetcher.FetchGolangciLintVersion()
	if err != nil {
		return fmt.Errorf("fetch golangci lint version: %w", err)
	}

	tmpl := templates.Template{
		ProjectName:         projectName,
		ModuleName:          moduleName,
		GoVersion:           goVersion,
		GolangciLintVersion: "v" + golangciLintVersion,
	}

	files := []struct {
		path    string
		content []byte
	}{
		{
			path:    fmt.Sprintf("%s/workspace/%s/cmd/%s/main.go", userHome, projectName, projectName),
			content: tmpl.MainGo(),
		},
		{
			path:    fmt.Sprintf("%s/workspace/%s/internal/app/flow/service.go", userHome, projectName),
			content: tmpl.ServiceGo(),
		},
		{
			path:    fmt.Sprintf("%s/workspace/%s/.gitignore", userHome, projectName),
			content: tmpl.Gitignore(),
		},
		{
			path:    fmt.Sprintf("%s/workspace/%s/.golangci.yaml", userHome, projectName),
			content: tmpl.Golangci(),
		},
		{
			path:    fmt.Sprintf("%s/workspace/%s/.pre-commit-config.yaml", userHome, projectName),
			content: tmpl.PreCommitConfig(),
		},
		{
			path:    fmt.Sprintf("%s/workspace/%s/go.mod", userHome, projectName),
			content: tmpl.GoMod(),
		},
		{
			path:    fmt.Sprintf("%s/workspace/%s/Makefile", userHome, projectName),
			content: tmpl.Makefile(),
		},
	}

	for _, file := range files {
		err := s.filesystem.CreateFile(file.path, file.content)
		if err != nil {
			return fmt.Errorf("create file %s: %w", file.path, err)
		}
	}

	return nil
}
