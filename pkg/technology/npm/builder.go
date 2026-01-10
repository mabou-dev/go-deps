package npm

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/mabou-dev/go-deps/pkg/logger"
	"github.com/mabou-dev/go-deps/pkg/model"
)

type NpmBuilder struct {
	Registry model.Registry

	Logger logger.Logger
}

func NewNpmBuilder() *NpmBuilder {
	return &NpmBuilder{}
}

func (n *NpmBuilder) CanHandle(projectPath string) bool {
	packageJsonPath := filepath.Join(projectPath, PACKAGE_LOCK_FILE)
	if _, err := os.Stat(packageJsonPath); err == nil {
		return true
	}
	return false
}

func (n *NpmBuilder) GetName() string {
	return NAME
}

func (n *NpmBuilder) BuildTree(projectPath string) (*model.NodeDependency, error) {
	// Implementation for building the dependency tree from package.json and package-lock.json
	//return &model.NodeDependency{
	//	Name:    "example-package",
	//	Version: "1.0.0",
	//	URLs:    []string{"https://registry.npmjs.org/example-package/-/example-package-1.0.0.tgz"},
	//}, nil
	err := errors.New("NpmBuilder BuildTree not implemented")
	n.Logger.Error(err.Error())
	return nil, err
}

func (n *NpmBuilder) DownloadPackage(pkg *model.NodeDependency, output string) error {
	return n.Registry.DownloadPackage(pkg, output)
}
