// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package tailcfg aliases canonical Tailscale protocol types to this fork.
package tailcfg

import forktailcfg "github.com/Xinlong-Wu/tailscale-oh/tailcfg"

type (
	// Node is the fork's node description type.
	Node = forktailcfg.Node
	// PeerCapability is the fork's peer capability identifier.
	PeerCapability = forktailcfg.PeerCapability
	// PeerCapMap is the fork's peer capability map.
	PeerCapMap = forktailcfg.PeerCapMap
	// RawMessage is the fork's raw JSON capability value.
	RawMessage = forktailcfg.RawMessage
	// UserProfile is the fork's user profile type.
	UserProfile = forktailcfg.UserProfile
)

// UnmarshalCapJSON decodes capability values using the fork implementation.
func UnmarshalCapJSON[T any](cm PeerCapMap, cap PeerCapability) ([]T, error) {
	return forktailcfg.UnmarshalCapJSON[T](cm, cap)
}
