package util

import (
	"os"
)

// OSInterface collects system level operations that need to be mocked out during tests.
type OSInterface interface {
	MkdirAll(path string, perm os.FileMode) error
}

// RealOS is used to dispatch the real system level operations.
type RealOS struct{}

// MkdirAll will call os.MkdirAll to create a directory.
func (RealOS) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}
