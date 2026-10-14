// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package controlhttpserver forwards the canonical control HTTP server path.
package controlhttpserver

import forkcontrolhttpserver "github.com/Xinlong-Wu/tailscale-oh/control/controlhttp/controlhttpserver"

var AcceptHTTP = forkcontrolhttpserver.AcceptHTTP
