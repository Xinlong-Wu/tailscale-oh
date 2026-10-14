// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package set aliases canonical set types to this fork.
package set

import forkset "github.com/Xinlong-Wu/tailscale-oh/util/set"

type Set[T comparable] = forkset.Set[T]
