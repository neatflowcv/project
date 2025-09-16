package realversionfetcher

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	versionfetcher "github.com/neatflowcv/project/internal/pkg/versionfetcher/core"
)

var _ versionfetcher.VersionFetcher = (*RealVersionFetcher)(nil)

var (
	ErrUnexpectedStatus = errors.New("unexpected status from GitHub API")
	ErrVersionMissing   = errors.New("version not found in response")
)

type RealVersionFetcher struct{}

func NewRealVersionFetcher() *RealVersionFetcher {
	return &RealVersionFetcher{}
}

func (f *RealVersionFetcher) FetchGoVersion(ctx context.Context) (string, error) {
	const endpoint = "https://go.dev/dl/?mode=json"

	content, err := f.fetch(ctx, endpoint)
	if err != nil {
		return "", fmt.Errorf("fetch go version: %w", err)
	}

	var payload []struct {
		Version string `json:"version"`
	}

	err = json.Unmarshal(content, &payload)
	if err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}

	if len(payload) == 0 || payload[0].Version == "" {
		return "", ErrVersionMissing
	}

	version := strings.TrimPrefix(payload[0].Version, "go")

	return version, nil
}

func (f *RealVersionFetcher) FetchGolangciLintVersion(ctx context.Context) (string, error) {
	const endpoint = "https://api.github.com/repos/golangci/golangci-lint/releases/latest"

	content, err := f.fetch(ctx, endpoint)
	if err != nil {
		return "", fmt.Errorf("fetch golangci lint version: %w", err)
	}

	var payload struct {
		TagName string `json:"tag_name"`
	}

	err = json.Unmarshal(content, &payload)
	if err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}

	if payload.TagName == "" {
		return "", ErrVersionMissing
	}

	version := payload.TagName
	version = strings.TrimPrefix(version, "v")

	return version, nil
}

func (*RealVersionFetcher) fetch(ctx context.Context, endpoint string) ([]byte, error) {
	const (
		requestTimeout = 15 * time.Second
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("User-Agent", "neatflowcv-project/versionfetcher")

	client := &http.Client{Timeout: requestTimeout} //nolint:exhaustruct

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %w: %d", ErrUnexpectedStatus, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	return body, nil
}
