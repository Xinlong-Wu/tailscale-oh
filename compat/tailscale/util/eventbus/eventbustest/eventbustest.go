// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package eventbustest forwards canonical event-bus test helpers to this fork.
package eventbustest

import forkeventbustest "github.com/Xinlong-Wu/tailscale-oh/util/eventbus/eventbustest"

var NewBus = forkeventbustest.NewBus
