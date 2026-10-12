// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package tsdial forwards the canonical dialer path to this fork.
package tsdial

import forktsdial "github.com/Xinlong-Wu/tailscale-oh/net/tsdial"

type Dialer = forktsdial.Dialer

var NewDialer = forktsdial.NewDialer
