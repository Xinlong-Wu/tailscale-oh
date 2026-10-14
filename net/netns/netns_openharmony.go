// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build openharmony

package netns

import (
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
	"syscall"

	"github.com/Xinlong-Wu/tailscale-oh/net/netmon"
	"github.com/Xinlong-Wu/tailscale-oh/types/logger"
	"golang.org/x/sys/unix"
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

// SetListenConfigInterfaceName sets lc.Control so that sockets created with
// it are bound to the named interface. OpenHarmony uses the Linux socket API
// for SO_BINDTODEVICE, and peerapi uses this to keep its kernel listener off
// physical interfaces.
//
// As on Linux, failures to set the socket option are logged and otherwise
// ignored so that a missing capability does not take out the listener.
func SetListenConfigInterfaceName(lc *net.ListenConfig, ifName string) error {
	if lc == nil {
		return errors.New("nil ListenConfig")
	}
	if lc.Control != nil {
		return errors.New("ListenConfig.Control already set")
	}
	lc.Control = func(network, address string, c syscall.RawConn) error {
		var sockErr error
		err := c.Control(func(fd uintptr) {
			sockErr = unix.SetsockoptString(int(fd), unix.SOL_SOCKET, unix.SO_BINDTODEVICE, ifName)
		})
		if sockErr != nil {
			log.Printf("[unexpected] netns: SO_BINDTODEVICE(%q) on %v %v: %v", ifName, network, address, sockErr)
		}
		if err != nil {
			return fmt.Errorf("RawConn.Control on %T: %w", c, err)
		}
		return nil
	}
	return nil
}
