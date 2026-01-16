// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !linux && !windows && !darwin && !openbsd

package netns

import (
	"syscall"

	"github.com/Xinlong-Wu/tailscale-oh/net/netmon"
	"github.com/Xinlong-Wu/tailscale-oh/types/logger"
)

func control(logger.Logf, *netmon.Monitor) func(network, address string, c syscall.RawConn) error {
	return controlC
}

// controlC does nothing to c.
func controlC(network, address string, c syscall.RawConn) error {
	return nil
}

func UseSocketMark() bool {
	return false
}
