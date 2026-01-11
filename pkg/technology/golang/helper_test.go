package golang

import (
	"path/filepath"
	"runtime"
	"testing"
)

// HelperTestDataPath returns the path to the test golang folder (e.g., <repo>/test/golang).
func HelperTestDataPath(t *testing.T) string {
	t.Helper()
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return filepath.Join("test", "golang")
	}
	dir := filepath.Dir(filename)
	return filepath.Clean(filepath.Join(dir, "..", "..", "..", "test", "golang"))
}
