// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package tsaddr forwards canonical Tailscale address helpers to this fork.
package tsaddr

import forktsaddr "github.com/Xinlong-Wu/tailscale-oh/net/tsaddr"

var (
	AllIPv4       = forktsaddr.AllIPv4
	AllIPv6       = forktsaddr.AllIPv6
	Tailscale4To6 = forktsaddr.Tailscale4To6
)
