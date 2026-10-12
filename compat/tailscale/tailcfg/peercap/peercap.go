// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package peercap aliases canonical peer capabilities to this fork.
package peercap

import forkpeercap "github.com/Xinlong-Wu/tailscale-oh/tailcfg/peercap"

type Cap = forkpeercap.Cap

const (
	Ingress     = forkpeercap.Ingress
	Relay       = forkpeercap.Relay
	RelayTarget = forkpeercap.RelayTarget
)
