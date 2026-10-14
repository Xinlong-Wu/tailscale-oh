// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package must forwards canonical must-succeed helpers to this fork.
package must

import forkmust "github.com/Xinlong-Wu/tailscale-oh/util/must"

func Do(err error) {
	forkmust.Do(err)
}

func Get[T any](v T, err error) T {
	return forkmust.Get(v, err)
}
