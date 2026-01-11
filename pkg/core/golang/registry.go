package golang

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/mabou-dev/go-deps/pkg/core"
	"github.com/mabou-dev/go-deps/pkg/utils/logger"
)

type GolangRegistry struct {
	BaseURL string

	Logger logger.Logger
}

func NewGolangRegistry(logger logger.Logger, baseURL string) *GolangRegistry {
	return &GolangRegistry{
		BaseURL: baseURL,
		Logger:  logger,
	}
}

func (r *GolangRegistry) DownloadPackage(pkg *core.NodeDependency, output string) error {
	err := errors.New("GolangRegistry DownloadPackage not implemented")
	r.Logger.Error(err.Error())
	return err
}

func (r *GolangRegistry) GetPackageURL(name string, version string) string {
	return fmt.Sprintf("%s/%s/@v/%s.zip", r.BaseURL, name, version)
}

func (r *GolangRegistry) GetMetadataURL(name string, version string) string {
	return fmt.Sprintf("%s/%s/@v/%s.mod", r.BaseURL, name, version)
}

func (r *GolangRegistry) FetchPackageMetadata(pkg *core.NodeDependency) ([]byte, error) {
	url := r.GetMetadataURL(pkg.Name, pkg.Version)
	resp, err := http.Get(url)
	if err != nil {
		r.Logger.Error("Failed to fetch package metadata: " + err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("failed to fetch package metadata, status code: %d", resp.StatusCode)
		r.Logger.Error(err.Error())
		return nil, err
	}

	return io.ReadAll(resp.Body)
}
