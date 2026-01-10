package npm

import (
	"errors"

	"github.com/mabou-dev/go-deps/pkg/logger"
	"github.com/mabou-dev/go-deps/pkg/model"
)

type NpmRegistry struct {
	BaseURL string

	Logger logger.Logger
}

func NewNpmRegistry(baseURL string) *NpmRegistry {
	return &NpmRegistry{
		BaseURL: baseURL,
	}
}

func (r *NpmRegistry) DownloadPackage(pkg *model.NodeDependency, output string) error {
	err := errors.New("NpmRegistry DownloadPackage not implemented")
	r.Logger.Error(err.Error())
	return err
}
