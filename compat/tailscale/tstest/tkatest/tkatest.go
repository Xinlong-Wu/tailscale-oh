// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package tkatest forwards canonical Tailnet Lock test helpers to this fork.
package tkatest

import forktkatest "github.com/Xinlong-Wu/tailscale-oh/tstest/tkatest"

var (
	HandleTKABootstrap  = forktkatest.HandleTKABootstrap
	HandleTKAInitBegin  = forktkatest.HandleTKAInitBegin
	HandleTKAInitFinish = forktkatest.HandleTKAInitFinish
	HandleTKASign       = forktkatest.HandleTKASign
	HandleTKASyncOffer  = forktkatest.HandleTKASyncOffer
)
