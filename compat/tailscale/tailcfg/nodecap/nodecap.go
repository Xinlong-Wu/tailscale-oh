// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package nodecap aliases canonical node capabilities to this fork.
package nodecap

import forknodecap "github.com/Xinlong-Wu/tailscale-oh/tailcfg/nodecap"

type Cap = forknodecap.Cap

const (
	DisableUPnP = forknodecap.DisableUPnP
	FileSharing = forknodecap.FileSharing
	Funnel      = forknodecap.Funnel
	FunnelPorts = forknodecap.FunnelPorts
	HTTPS       = forknodecap.HTTPS
	SSH         = forknodecap.SSH
)
