package versionfetcher

import "context"

type VersionFetcher interface {
	FetchGoVersion(ctx context.Context) (string, error)           // ex> 1.25.1
	FetchGolangciLintVersion(ctx context.Context) (string, error) // ex> 2.4.0
}
