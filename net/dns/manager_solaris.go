// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

package dns

import (
	"github.com/Xinlong-Wu/tailscale-oh/control/controlknobs"
	"github.com/Xinlong-Wu/tailscale-oh/health"
	"github.com/Xinlong-Wu/tailscale-oh/types/logger"
	"github.com/Xinlong-Wu/tailscale-oh/util/eventbus"
	"github.com/Xinlong-Wu/tailscale-oh/util/syspolicy/policyclient"
)

func NewOSConfigurator(logf logger.Logf, health *health.Tracker, bus *eventbus.Bus, _ policyclient.Client, _ *controlknobs.Knobs, iface string) (OSConfigurator, error) {
	return newDirectManager(logf, health, bus), nil
}
