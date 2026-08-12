// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package metrics aliases canonical Tailscale metrics types to this fork.
package metrics

import forkmetrics "github.com/Xinlong-Wu/tailscale-oh/metrics"

// Set is the fork's expvar metric set.
type Set = forkmetrics.Set

// LabelMap is the fork's labelled metric map.
type LabelMap = forkmetrics.LabelMap
