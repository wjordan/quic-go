//go:build !linux

package self_test

func isPermissionError(err error) bool {
	return false
}

func pmtuDiscoverySupported() bool {
	return true
}
