// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"testing"

	"github.com/Xinlong-Wu/tailscale-oh/tstest/deptest"
)

func TestDeps(t *testing.T) {
	deptest.DepChecker{
		BadDeps: map[string]string{
			"github.com/Xinlong-Wu/tailscale-oh/tailcfg": "circular dependency via go generate",
			"github.com/Xinlong-Wu/tailscale-oh/version": "circular dependency via go generate",
		},
	}.Check(t)
}
