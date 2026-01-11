package npm

import (
	"os"
	"path/filepath"

	"github.com/mabou-dev/go-deps/pkg/core"
	"github.com/mabou-dev/go-deps/pkg/utils/logger"
)

type NpmTreeBuilder struct {
	Registry *NpmRegistry
	Logger   logger.Logger
}

func NewNpmTreeBuilder(logger logger.Logger, registry *NpmRegistry) *NpmTreeBuilder {
	return &NpmTreeBuilder{
		Registry: registry,
		Logger:   logger,
	}
}

func (n *NpmTreeBuilder) CanHandle(projectPath string) bool {
	packageJsonPath := filepath.Join(projectPath, PACKAGE_LOCK_FILE)
	if _, err := os.Stat(packageJsonPath); err == nil {
		return true
	}
	return false
}

func (n *NpmTreeBuilder) GetName() string {
	return NAME
}

func (n *NpmTreeBuilder) BuildTree(projectPath string) (*core.NodeDependency, error) {
	filename := filepath.Join(projectPath, PACKAGE_LOCK_FILE)
	tree := n.buildTreeFromLockFile(filename)
	return tree, nil
}

func (n *NpmTreeBuilder) buildTreeFromLockFile(filename string) *core.NodeDependency {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		n.Logger.Error("package-lock.json not found at path: " + filename)
		return &core.NodeDependency{}
	}

	lockJson, err := ParsePackageLockJSON(filename)
	if err != nil {
		n.Logger.Error("Failed to parse package-lock.json: " + err.Error())
		return &core.NodeDependency{}
	}

	tree := PackageLockJSONToNodeDependency(lockJson)
	return tree
}

func (n *NpmTreeBuilder) DownloadPackage(pkg *core.NodeDependency, output string) error {
	return n.Registry.DownloadPackage(pkg, output)
}
