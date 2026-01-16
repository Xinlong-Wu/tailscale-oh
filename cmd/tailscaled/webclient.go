// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !ts_omit_webclient

package main

import (
	"github.com/Xinlong-Wu/tailscale-oh/client/local"
	"github.com/Xinlong-Wu/tailscale-oh/ipn/ipnlocal"
	"github.com/Xinlong-Wu/tailscale-oh/paths"
)

func init() {
	hookConfigureWebClient.Set(func(lb *ipnlocal.LocalBackend) {
		lb.ConfigureWebClient(&local.Client{
			Socket:        args.socketpath,
			UseSocketOnly: args.socketpath != paths.DefaultTailscaledSocket(),
		})
	})
}
