// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package zstdframe forwards canonical zstd frame helpers to this fork.
package zstdframe

import forkzstdframe "github.com/Xinlong-Wu/tailscale-oh/util/zstdframe"

var (
	AppendEncode       = forkzstdframe.AppendEncode
	FastestCompression = forkzstdframe.FastestCompression
)
