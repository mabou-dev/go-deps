package golang

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanHandle(t *testing.T) {
	builder := NewGolangTreeBuilder(nil, nil)

	projectPath := HelperTestDataPath(t)
	canHandle := builder.CanHandle(projectPath)
	assert.True(t, canHandle, "GolangTreeBuilder should handle project with go.mod")

	invalidProjectPath := filepath.Join(HelperTestDataPath(t), "no-golang-project")
	canHandle = builder.CanHandle(invalidProjectPath)
	assert.False(t, canHandle, "GolangTreeBuilder should not handle project without go.mod")
}

func TestBuildTreeFromModFile(t *testing.T) {
	path := filepath.Join(HelperTestDataPath(t), GO_MOD_FILE)
	builder := NewGolangTreeBuilder(nil, nil)
	tree, err := builder.buildTreeFromModFile(path)
	assert.NoError(t, err)
	assert.NotNil(t, tree)
	assert.Equal(t, "github.com/example/project", tree.Name)
	assert.NotEmpty(t, tree.Dependencies)

	found := map[string]string{}
	for _, d := range tree.Dependencies {
		found[d.Name] = d.Version
	}

	assert.Contains(t, found, "github.com/spf13/cobra")
	assert.Equal(t, "v1.8.1", found["github.com/spf13/cobra"], "cobra version")

	assert.Contains(t, found, "github.com/spf13/viper")
	assert.Equal(t, "v1.19.0", found["github.com/spf13/viper"], "viper version")

	assert.Contains(t, found, "github.com/stretchr/testify")
	assert.Equal(t, "v1.9.0", found["github.com/stretchr/testify"], "testify version")

	assert.Contains(t, found, "go.uber.org/zap")
	assert.Equal(t, "v1.27.1", found["go.uber.org/zap"], "zap version")
}
