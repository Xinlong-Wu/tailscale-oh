// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package atomicfile forwards the canonical Tailscale package path to this
// fork. It exists for external modules that still import tailscale.com.
package atomicfile

import (
	"os"

	forkatomicfile "github.com/Xinlong-Wu/tailscale-oh/atomicfile"
)

// WriteFile forwards to the fork's atomicfile.WriteFile.
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	return forkatomicfile.WriteFile(filename, data, perm)
}

// Rename forwards to the fork's atomicfile.Rename.
func Rename(srcFile, dstFile string) error {
	return forkatomicfile.Rename(srcFile, dstFile)
}
