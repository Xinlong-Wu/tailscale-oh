// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !ts_omit_routecheck

package ipnlocal

import (
	"github.com/Xinlong-Wu/tailscale-oh/net/routecheck"
	"github.com/Xinlong-Wu/tailscale-oh/tailcfg"
)

func isRouteCheckEnabled(self tailcfg.NodeView) bool {
	return routecheck.IsEnabled(self)
}
