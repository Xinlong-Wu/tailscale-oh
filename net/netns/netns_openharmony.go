// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build openharmony

package netns

import (
	"fmt"
	"sync"
	"syscall"

	"github.com/Xinlong-Wu/tailscale-oh/net/netmon"
	"github.com/Xinlong-Wu/tailscale-oh/types/logger"
)

var (
	openharmonyProtectFuncMu sync.Mutex
	openharmonyProtectFunc   func(fd int) error
)

// UseSocketMark reports whether SO_MARK is in use. OpenHarmony uses the
// application-provided protect hook instead.
func UseSocketMark() bool {
	return false
}

// SetOpenHarmonyProtectFunc registers a function that protects a socket from
// being routed back through the VPN. A nil function disables the hook.
func SetOpenHarmonyProtectFunc(f func(fd int) error) {
	openharmonyProtectFuncMu.Lock()
	defer openharmonyProtectFuncMu.Unlock()
	openharmonyProtectFunc = f
}

func control(logger.Logf, *netmon.Monitor) func(network, address string, c syscall.RawConn) error {
	return controlC
}

// controlC marks c as necessary to dial in a separate network namespace.
//
// It's intentionally the same signature as net.Dialer.Control
// and net.ListenConfig.Control.
func controlC(network, address string, c syscall.RawConn) error {
	var sockErr error
	err := c.Control(func(fd uintptr) {
		openharmonyProtectFuncMu.Lock()
		f := openharmonyProtectFunc
		openharmonyProtectFuncMu.Unlock()
		if f != nil {
			sockErr = f(int(fd))
		}
	})
	if err != nil {
		return fmt.Errorf("RawConn.Control on %T: %w", c, err)
	}
	return sockErr
}
