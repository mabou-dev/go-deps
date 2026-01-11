package golang

import (
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"

	"github.com/mabou-dev/go-deps/pkg/core"
	"github.com/mabou-dev/go-deps/pkg/utils/logger"
)

type GolangTreeBuilder struct {
	Registry *GolangRegistry

	Logger logger.Logger
}

func NewGolangTreeBuilder(logger logger.Logger, registry *GolangRegistry) *GolangTreeBuilder {
	return &GolangTreeBuilder{
		Registry: registry,
		Logger:   logger,
	}
}

func (g *GolangTreeBuilder) CanHandle(projectPath string) bool {
	goModPath := filepath.Join(projectPath, GO_MOD_FILE)
	if _, err := os.Stat(goModPath); err == nil {
		return true
	}
	return false
}

func (g *GolangTreeBuilder) GetName() string {
	return NAME
}

func (g *GolangTreeBuilder) BuildTree(projectPath string) (*core.NodeDependency, error) {
	filename := filepath.Join(projectPath, GO_MOD_FILE)
	return g.buildTreeFromModFile(filename)
}

func (g *GolangTreeBuilder) buildTreeFromModFile(filename string) (*core.NodeDependency, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		g.Logger.Error("go.mod not found at path: " + filename)
		return &core.NodeDependency{}, err
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		g.Logger.Error("Failed to read go.mod: " + err.Error())
		return &core.NodeDependency{}, err
	}

	modFile, err := modfile.Parse(filename, data, nil)
	if err != nil {
		g.Logger.Error("Failed to parse go.mod: " + err.Error())
		return &core.NodeDependency{}, err
	}

	root := &core.NodeDependency{
		Name:    modFile.Module.Mod.Path,
		Version: "", // Go modules do not have a version in go.mod
	}

	for _, req := range modFile.Require {
		dep := &core.NodeDependency{
			Name:    req.Mod.Path,
			Version: req.Mod.Version,
		}
		root.Dependencies = append(root.Dependencies, dep)
	}

	return root, nil
}

func (g *GolangTreeBuilder) resolvePackage(depName, depVersion string) (*core.NodeDependency, error) {
	url := g.Registry.GetPackageURL(depName, depVersion)

	node := &core.NodeDependency{
		Name:    depName,
		Version: depVersion,
		URLs:    []string{url},
	}

	data, err := g.Registry.FetchPackageMetadata(node)
	if err != nil {
		g.Logger.Error("Failed to fetch package metadata for " + depName + ": " + err.Error())
		return nil, err
	}

	modFile, err := modfile.Parse(GO_MOD_FILE, data, nil)
	if err != nil {
		g.Logger.Error("Failed to parse module file for " + depName + ": " + err.Error())
		return nil, err
	}

	for _, req := range modFile.Require {
		child, err := g.resolvePackage(req.Mod.Path, req.Mod.Version)
		if err != nil {
			g.Logger.Error("Failed to resolve package " + req.Mod.Path + ": " + err.Error())
			return nil, err
		}
		node.Dependencies = append(node.Dependencies, child)
	}
	return node, nil
}

func (g *GolangTreeBuilder) DownloadPackage(pkg *core.NodeDependency, output string) error {
	return g.Registry.DownloadPackage(pkg, output)
}
