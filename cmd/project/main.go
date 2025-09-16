package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/neatflowcv/project/internal/app/flow"
	realfilesystem "github.com/neatflowcv/project/internal/pkg/filesystem/real"
	"github.com/urfave/cli/v3"
)

func main() {
	var project string

	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	service := flow.NewService(realfilesystem.NewRealFilesystem())

	cmd := cli.Command{ //nolint:exhaustruct
		Commands: []*cli.Command{
			{
				Name:    "new",
				Usage:   "new project",
				Aliases: []string{"n"},
				Action: func(ctx context.Context, c *cli.Command) error {
					if project == "" {
						return cli.Exit("project name is required", 1)
					}

					projectName := extractProjectName(project)
					moduleName := extractModuleName(project)
					log.Println("create project", projectName, moduleName)
					err := service.NewProject(userHome, projectName, moduleName)
					if err != nil {
						return cli.Exit(err.Error(), 1)
					}

					return nil
				},
				Arguments: []cli.Argument{
					&cli.StringArg{ //nolint:exhaustruct
						Name:        "name",
						Destination: &project,
						UsageText:   "https://github.com/neatflowcv/project or https://github.com/neatflowcv/project.git or project",
					},
				},
			},
		},
	}

	err = cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func extractModuleName(projectName string) string {
	ret := strings.TrimSuffix(projectName, ".git")
	ret = strings.TrimPrefix(ret, "https://")
	ret = strings.TrimPrefix(ret, "http://")

	return ret
}

func extractProjectName(project string) string {
	project = strings.TrimSuffix(project, ".git")
	parts := strings.Split(project, "/")
	ret := parts[len(parts)-1]

	return ret
}
