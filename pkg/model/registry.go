package model

type Registry interface {
	// DownloadPackage downloads the given package to the specified output location
	DownloadPackage(pkg *NodeDependency, output string) error
}
