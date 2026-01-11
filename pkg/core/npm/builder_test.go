package npm

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanHandle(t *testing.T) {
	builder := NewNpmTreeBuilder(nil, nil)

	projectPath := HelperTestDataPath(t)
	canHandle := builder.CanHandle(projectPath)
	assert.True(t, canHandle, "NpmTreeBuilder should handle project with package-lock.json")

	invalidProjectPath := filepath.Join(HelperTestDataPath(t), "no-npm-project")
	canHandle = builder.CanHandle(invalidProjectPath)
	assert.False(t, canHandle, "NpmTreeBuilder should not handle project without package-lock.json")
}

func TestBuildTreeFromLockFile(t *testing.T) {
	path := filepath.Join(HelperTestDataPath(t), PACKAGE_LOCK_FILE)
	builder := NewNpmTreeBuilder(nil, nil)
	tree := builder.buildTreeFromLockFile(path)
	assert.NotNil(t, tree)
	assert.Equal(t, "example-project", tree.Name)
	assert.NotEmpty(t, tree.Dependencies)

	found := map[string]string{}
	for _, d := range tree.Dependencies {
		found[d.Name] = d.Version
	}

	assert.Contains(t, found, "left-pad")
	assert.Equal(t, "1.3.0", found["left-pad"], "left-pad version")

	assert.Contains(t, found, "is-odd")
	assert.Equal(t, "3.0.1", found["is-odd"], "is-odd version")
}
