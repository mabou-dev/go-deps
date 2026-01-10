package npm

import (
	"path/filepath"
	"runtime"
	"testing"
)

// HelperTestDataPath returns the path to the test npm folder (e.g., <repo>/test/npm).
// It's in a _test.go file so it's only available to tests in this package.
func HelperTestDataPath(t *testing.T) string {
	t.Helper()
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return filepath.Join("test", "npm")
	}
	dir := filepath.Dir(filename)
	return filepath.Clean(filepath.Join(dir, "..", "..", "..", "test", "npm"))
}
