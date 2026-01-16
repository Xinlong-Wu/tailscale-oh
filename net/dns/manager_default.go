// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build (!linux || android) && !freebsd && !openbsd && !windows && !darwin && !illumos && !solaris && !plan9

package dns

import (
	"github.com/Xinlong-Wu/tailscale-oh/control/controlknobs"
	"github.com/Xinlong-Wu/tailscale-oh/health"
	"github.com/Xinlong-Wu/tailscale-oh/types/logger"
	"github.com/Xinlong-Wu/tailscale-oh/util/eventbus"
	"github.com/Xinlong-Wu/tailscale-oh/util/syspolicy/policyclient"
)

// NewOSConfigurator creates a new OS configurator.
//
// The health tracker and the knobs may be nil and are ignored on this platform.
func NewOSConfigurator(logger.Logf, *health.Tracker, *eventbus.Bus, policyclient.Client, *controlknobs.Knobs, string) (OSConfigurator, error) {
	return NewNoopManager()
}
