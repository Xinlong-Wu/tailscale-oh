// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package key aliases canonical key types to this fork.
package key

import forkkey "github.com/Xinlong-Wu/tailscale-oh/types/key"

type (
	ControlPrivate = forkkey.ControlPrivate
	DiscoPublic    = forkkey.DiscoPublic
	MachinePrivate = forkkey.MachinePrivate
	MachinePublic  = forkkey.MachinePublic
	NodePublic     = forkkey.NodePublic
)

var (
	NewControl = forkkey.NewControl
	NewDisco   = forkkey.NewDisco
	NewMachine = forkkey.NewMachine
	NewNode    = forkkey.NewNode
)
