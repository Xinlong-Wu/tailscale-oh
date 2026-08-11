// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

package tailscaleroot

import (
	"os"
	"testing"

	"golang.org/x/mod/modfile"
)

func TestGoMod(t *testing.T) {
	goMod, err := os.ReadFile("go.mod")
	if err != nil {
		t.Fatal(err)
	}
	f, err := modfile.Parse("go.mod", goMod, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(f.Replace) != 1 {
		t.Fatalf("go.mod has %d replace directives; expect the setec compatibility replacement", len(f.Replace))
	}
	r := f.Replace[0]
	if r.Old.Path != "tailscale.com" || r.Old.Version != "" ||
		r.New.Path != "./compat/tailscale" || r.New.Version != "" {
		t.Errorf("unexpected go.mod replace directive: %s %s => %s %s", r.Old.Path, r.Old.Version, r.New.Path, r.New.Version)
	}
}
