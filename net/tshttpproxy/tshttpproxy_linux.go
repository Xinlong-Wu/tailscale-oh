// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build linux

package tshttpproxy

import (
	"net/http"
	"net/url"

	"github.com/Xinlong-Wu/tailscale-oh/feature/buildfeatures"
	"github.com/Xinlong-Wu/tailscale-oh/version/distro"
)

func init() {
	sysProxyFromEnv = linuxSysProxyFromEnv
}

func linuxSysProxyFromEnv(req *http.Request) (*url.URL, error) {
	if buildfeatures.HasSynology && distro.Get() == distro.Synology {
		return synologyProxyFromConfigCached(req)
	}
	return nil, nil
}
