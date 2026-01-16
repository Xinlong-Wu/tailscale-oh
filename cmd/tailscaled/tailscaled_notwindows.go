// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !windows && go1.19

package main // import "github.com/Xinlong-Wu/tailscale-oh/cmd/tailscaled"

import "github.com/Xinlong-Wu/tailscale-oh/logpolicy"

func isWindowsService() bool { return false }

func runWindowsService(pol *logpolicy.Policy) error { panic("unreachable") }

func beWindowsSubprocess() bool { return false }
