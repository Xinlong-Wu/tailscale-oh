// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package tka aliases canonical Tailnet Lock types to this fork.
package tka

import forktka "github.com/Xinlong-Wu/tailscale-oh/tka"

type (
	AUM              = forktka.AUM
	Authority        = forktka.Authority
	CompactableChonk = forktka.CompactableChonk
)

var (
	ChonkMem = forktka.ChonkMem
	Open     = forktka.Open
)
