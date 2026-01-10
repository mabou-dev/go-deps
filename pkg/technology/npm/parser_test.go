package npm

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePackageLockJSON(t *testing.T) {
	path := filepath.Join(HelperTestDataPath(t), PACKAGE_LOCK_FILE)
	pkg, err := ParsePackageLockJSON(path)
	assert.NoError(t, err)

	root := PackageLockJSONToNodeDependency(pkg)
	assert.NotNil(t, root)
	assert.NotEmpty(t, root.Dependencies)

	found := map[string]string{}
	for _, d := range root.Dependencies {
		found[d.Name] = d.Version
	}

	assert.Contains(t, found, "left-pad")
	assert.Equal(t, "1.3.0", found["left-pad"], "left-pad version")

	assert.Contains(t, found, "is-odd")
	assert.Equal(t, "3.0.1", found["is-odd"], "is-odd version")
}
