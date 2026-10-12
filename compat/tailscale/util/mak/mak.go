// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package mak forwards canonical map helpers to this fork.
package mak

import forkmak "github.com/Xinlong-Wu/tailscale-oh/util/mak"

func Set[K comparable, V any, T ~map[K]V](m *T, k K, v V) {
	forkmak.Set(m, k, v)
}
