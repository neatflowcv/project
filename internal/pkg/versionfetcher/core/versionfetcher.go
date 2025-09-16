package versionfetcher

type VersionFetcher interface {
	FetchGoVersion() (string, error) // ex> 1.25.1
	FetchGolangciLintVersion() (string, error) // ex> 2.4.0
}
