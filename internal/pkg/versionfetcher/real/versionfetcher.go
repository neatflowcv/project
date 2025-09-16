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
	ErrTagNameMissing   = errors.New("tag_name not found in response")
)

var (
	ErrVersionMissing = errors.New("version not found in response")
)

type RealVersionFetcher struct {
}

func NewRealVersionFetcher() *RealVersionFetcher {
	return &RealVersionFetcher{}
}

func (f *RealVersionFetcher) FetchGoVersion(ctx context.Context) (string, error) {
	const endpoint = "https://go.dev/dl/?mode=json"

	const (
		requestTimeout = 5 * time.Second
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("User-Agent", "neatflowcv-project/versionfetcher")

	client := &http.Client{Timeout: requestTimeout} //nolint:exhaustruct

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request go dl json: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)

		return "", fmt.Errorf("%w: status %d: %s", ErrUnexpectedStatus, resp.StatusCode, strings.TrimSpace(string(body)))
	}

	var payload []struct {
		Version string `json:"version"`
	}

	dec := json.NewDecoder(resp.Body)

	err = dec.Decode(&payload)
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

	const (
		requestTimeout = 5 * time.Second
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}
	// GitHub API recommends setting a User-Agent
	req.Header.Set("User-Agent", "neatflowcv-project/versionfetcher")

	client := &http.Client{Timeout: requestTimeout} //nolint:exhaustruct

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request latest release: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)

		return "", fmt.Errorf("%w: status %d: %s", ErrUnexpectedStatus, resp.StatusCode, strings.TrimSpace(string(body)))
	}

	var payload struct {
		TagName string `json:"tag_name"`
	}

	dec := json.NewDecoder(resp.Body)

	err = dec.Decode(&payload)
	if err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}

	if payload.TagName == "" {
		return "", fmt.Errorf("%w", ErrTagNameMissing)
	}

	version := payload.TagName
	// Our templates add a leading 'v', so return the numeric part if needed.
	version = strings.TrimPrefix(version, "v")

	return version, nil
}
