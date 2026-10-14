// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package syncs aliases canonical synchronization helpers to this fork.
package syncs

import forksyncs "github.com/Xinlong-Wu/tailscale-oh/syncs"

type Map[K comparable, V any] = forksyncs.Map[K, V]
