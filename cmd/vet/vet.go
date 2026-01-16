// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package vet is a tool to statically check Go source code.
package main

import (
	_ "embed"

	"golang.org/x/tools/go/analysis/unitchecker"
	"github.com/Xinlong-Wu/tailscale-oh/cmd/vet/jsontags"
	"github.com/Xinlong-Wu/tailscale-oh/cmd/vet/lowerell"
	"github.com/Xinlong-Wu/tailscale-oh/cmd/vet/subtestnames"
)

//go:embed jsontags_allowlist
var jsontagsAllowlistSource string

func init() {
	jsontags.RegisterAllowlist(jsontags.ParseAllowlist(jsontagsAllowlistSource))
	jsontags.RegisterPureIsZeroMethods(jsontags.PureIsZeroMethodsInTailscaleModule)
}

func main() {
	unitchecker.Main(jsontags.Analyzer, lowerell.Analyzer, subtestnames.Analyzer)
}
