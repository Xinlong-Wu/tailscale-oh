// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package logger aliases the canonical Tailscale logger type to this fork.
// It exists for external modules that still import tailscale.com.
package logger

import forklogger "github.com/Xinlong-Wu/tailscale-oh/types/logger"

// Logf is the fork's logging function type.
type Logf = forklogger.Logf

// Discard drops all log messages.
var Discard = forklogger.Discard
