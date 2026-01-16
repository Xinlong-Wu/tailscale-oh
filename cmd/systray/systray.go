// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build cgo || !darwin

// systray is a minimal Tailscale systray application.
package main

import (
	"flag"

	"github.com/Xinlong-Wu/tailscale-oh/client/local"
	"github.com/Xinlong-Wu/tailscale-oh/client/systray"
	"github.com/Xinlong-Wu/tailscale-oh/paths"
)

var socket = flag.String("socket", paths.DefaultTailscaledSocket(), "path to tailscaled socket")

func main() {
	flag.Parse()
	lc := &local.Client{Socket: *socket}
	new(systray.Menu).Run(lc)
}
