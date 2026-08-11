// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build for_go_mod_tidy_only

package gokrazydeps

import (
	_ "github.com/Xinlong-Wu/tailscale-oh/cmd/tailscale"
	_ "github.com/Xinlong-Wu/tailscale-oh/cmd/tailscaled"
	_ "github.com/Xinlong-Wu/tailscale-oh/cmd/tta"
	_ "github.com/gokrazy/gokrazy/cmd/dhcp"
	_ "github.com/gokrazy/kernel.arm64"
	_ "github.com/gokrazy/serial-busybox"
	_ "github.com/tailscale/ts-gokrazy/gokrazyinit"
)
