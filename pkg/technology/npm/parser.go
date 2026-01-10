package npm

import (
	"encoding/json"
	"os"

	"github.com/mabou-dev/go-deps/pkg/model"
)

type PackageJSON struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

func ParsePackageJSON(path string) (*PackageJSON, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var pkg PackageJSON
	if err = json.Unmarshal(data, &pkg); err != nil {
		return nil, err
	}

	return &pkg, nil
}

func PackageJSONToNodeDependency(pkg *PackageJSON) *model.NodeDependency {
	root := &model.NodeDependency{
		Name:    pkg.Name,
		Version: pkg.Version,
	}

	for name, version := range pkg.Dependencies {
		root.Dependencies = append(root.Dependencies, &model.NodeDependency{
			Name:    name,
			Version: version,
		})
	}
	for name, version := range pkg.DevDependencies {
		root.Dependencies = append(root.Dependencies, &model.NodeDependency{
			Name:    name,
			Version: version,
		})
	}

	return root
}

type PackageLockJSON struct {
	Name         string                    `json:"name"`
	Version      string                    `json:"version"`
	Dependencies map[string]lockDependency `json:"dependencies"`
}

type lockDependency struct {
	Version      string                    `json:"version"`
	Requires     map[string]string         `json:"requires"`
	Dependencies map[string]lockDependency `json:"dependencies"`
}

func ParsePackageLockJSON(path string) (*PackageLockJSON, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var pkg PackageLockJSON
	if err = json.Unmarshal(data, &pkg); err != nil {
		return nil, err
	}

	return &pkg, nil
}

func PackageLockJSONToNodeDependency(pkg *PackageLockJSON) *model.NodeDependency {
	root := &model.NodeDependency{
		Name:    pkg.Name,
		Version: pkg.Version,
	}

	for name, lockDep := range pkg.Dependencies {
		child := lockDependencyToNodeDependency(name, lockDep)
		root.Dependencies = append(root.Dependencies, child)
	}

	return root
}

func lockDependencyToNodeDependency(name string, dep lockDependency) *model.NodeDependency {
	node := &model.NodeDependency{
		Name:    name,
		Version: dep.Version,
	}
	for depName, depLock := range dep.Dependencies {
		child := lockDependencyToNodeDependency(depName, depLock)
		node.Dependencies = append(node.Dependencies, child)
	}
	return node
}
