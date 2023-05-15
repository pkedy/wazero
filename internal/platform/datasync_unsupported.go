//go:build !linux

package platform

import (
	"os"
	"syscall"
)

func datasync(f *os.File) syscall.Errno {
	// Attempt to sync everything, even if we only need to sync the data.
	return sync(f)
}
