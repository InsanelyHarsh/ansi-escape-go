package main

import (
	"syscall"
	"time"
	"unsafe"
)

func enableRawMode(fd int) (*syscall.Termios, error) {
	var old syscall.Termios

	if _, _, err := syscall.Syscall6(
		syscall.SYS_IOCTL,
		uintptr(fd),
		uintptr(syscall.TCGETS),
		uintptr(unsafe.Pointer(&old)),
		0, 0, 0,
	); err != 0 {
		return nil, err
	}

	raw := old

	// Input flags
	raw.Iflag &^= syscall.IGNBRK |
		syscall.BRKINT |
		syscall.PARMRK |
		syscall.ISTRIP |
		syscall.INLCR |
		syscall.IGNCR |
		syscall.ICRNL |
		syscall.IXON

	// Output flags
	raw.Oflag &^= syscall.OPOST

	// Local flags
	raw.Lflag &^= syscall.ECHO |
		syscall.ECHONL |
		syscall.ICANON |
		syscall.ISIG |
		syscall.IEXTEN

	// Character size
	raw.Cflag &^= syscall.CSIZE | syscall.PARENB
	raw.Cflag |= syscall.CS8

	// Read immediately, don't wait for Enter.
	raw.Cc[syscall.VMIN] = 1
	raw.Cc[syscall.VTIME] = 0

	if _, _, err := syscall.Syscall6(
		syscall.SYS_IOCTL,
		uintptr(fd),
		uintptr(syscall.TCSETS),
		uintptr(unsafe.Pointer(&raw)),
		0, 0, 0,
	); err != 0 {
		return nil, err
	}

	return &old, nil
}

// inputReady reports whether a byte is already available to read on fd
// within timeout, without blocking past it. Used to tell a standalone ESC
// key press (no follow-up byte) apart from the start of an escape sequence
// like ESC [ A, since VMIN=1/VTIME=0 makes a plain Read block indefinitely.
func inputReady(fd int, timeout time.Duration) bool {
	var fds syscall.FdSet
	fds.Bits[fd/64] |= 1 << (uint(fd) % 64)

	tv := syscall.NsecToTimeval(timeout.Nanoseconds())

	n, err := syscall.Select(fd+1, &fds, nil, nil, &tv)
	if err != nil {
		return false
	}

	return n > 0
}

func restore(fd int, old *syscall.Termios) {
	syscall.Syscall6(
		syscall.SYS_IOCTL,
		uintptr(fd),
		uintptr(syscall.TCSETS),
		uintptr(unsafe.Pointer(old)),
		0, 0, 0,
	)
}
