// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package multierr aliases the canonical Tailscale error helpers to this fork.
package multierr

import forkmultierr "github.com/Xinlong-Wu/tailscale-oh/util/multierr"

// Error is the fork's multiple-error type.
type Error = forkmultierr.Error

var (
	// New combines multiple errors.
	New = forkmultierr.New
	// Range walks an error tree.
	Range = forkmultierr.Range
)
