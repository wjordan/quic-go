//go:build linux

package self_test

import (
	"errors"
	"os"

	"golang.org/x/sys/unix"
)

// The first sendmsg call on a new UDP socket sometimes errors on Linux.
// It's not clear why this happens.
// See https://github.com/golang/go/issues/63322.
func isPermissionError(err error) bool {
	var serr *os.SyscallError
	if errors.As(err, &serr) {
		return serr.Syscall == "sendmsg" && serr.Err == unix.EPERM
	}
	return false
}

// pmtuDiscoverySupported probes whether the kernel supports setting the DF bit
// via IP_MTU_DISCOVER / IPV6_MTU_DISCOVER. In container environments (e.g. gVisor)
// these setsockopt calls may not be implemented.
func pmtuDiscoverySupported() bool {
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM|unix.SOCK_CLOEXEC, 0)
	if err != nil {
		return false
	}
	defer unix.Close(fd)
	return unix.SetsockoptInt(fd, unix.IPPROTO_IP, unix.IP_MTU_DISCOVER, unix.IP_PMTUDISC_PROBE) == nil
}
