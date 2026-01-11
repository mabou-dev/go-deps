package adapter

import (
	"fmt"

	"github.com/mabou-dev/go-deps/pkg/core"
	"github.com/mabou-dev/go-deps/pkg/core/golang"
	"github.com/mabou-dev/go-deps/pkg/core/npm"
	"github.com/mabou-dev/go-deps/pkg/utils/config"
	"github.com/mabou-dev/go-deps/pkg/utils/logger"
)

type Adapter interface {
	CanHandle(projectPath string) bool
	GetName() string

	core.DependencyTreeBuilder
	core.Registry
}

var registry = []Adapter{}

func RegisterAdapter(adapter Adapter) {
	registry = append(registry, adapter)
}

func Resolve(projectPath string) (Adapter, error) {
	for _, adapter := range registry {
		if adapter.CanHandle(projectPath) {
			return adapter, nil
		}
	}
	return nil, fmt.Errorf("no adapter found for project path: %s", projectPath)
}

func Init(logger logger.Logger, cfg map[string]config.TechnicalConfig) {
	// Initialize and register all available adapters here
	var _ Adapter = (*npm.NpmTreeBuilder)(nil)
	NpmRegistry := npm.NewNpmRegistry(logger, cfg[npm.NAME].RegistryURL)
	RegisterAdapter(npm.NewNpmTreeBuilder(logger, NpmRegistry))

	var _ Adapter = (*golang.GolangTreeBuilder)(nil)
	GolangRegistry := golang.NewGolangRegistry(logger, cfg[golang.NAME].RegistryURL)
	RegisterAdapter(golang.NewGolangTreeBuilder(logger, GolangRegistry))
}
