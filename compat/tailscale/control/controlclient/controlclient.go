// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package controlclient aliases the canonical control client to this fork.
package controlclient

import forkcontrolclient "github.com/Xinlong-Wu/tailscale-oh/control/controlclient"

type (
	// Direct is the fork's direct control client.
	Direct = forkcontrolclient.Direct
	// Options configures a control client.
	Options = forkcontrolclient.Options
)

const LoginDefault = forkcontrolclient.LoginDefault

var NewDirect = forkcontrolclient.NewDirect
