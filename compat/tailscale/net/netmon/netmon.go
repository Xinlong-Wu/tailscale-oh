// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package netmon forwards canonical network-monitor helpers to this fork.
package netmon

import forknetmon "github.com/Xinlong-Wu/tailscale-oh/net/netmon"

type Monitor = forknetmon.Monitor

var NewStatic = forknetmon.NewStatic
