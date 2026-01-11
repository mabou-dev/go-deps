package core

type NodeDependency struct {
	Name         string            `json:"name"`
	Version      string            `json:"version"`
	URLs         []string          `json:"urls"`
	Dependencies []*NodeDependency `json:"dependencies"`
}

type DependencyTreeBuilder interface {
	// BuildTree constructs the dependency tree for the project at the given path
	BuildTree(projectPath string) (*NodeDependency, error)
}

type Registry interface {
	// DownloadPackage downloads the given package to the specified output location
	DownloadPackage(pkg *NodeDependency, output string) error
}
