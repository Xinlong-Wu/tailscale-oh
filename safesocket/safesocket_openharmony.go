// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build openharmony

package safesocket

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func init() {
	localTCPPortAndToken = localTCPPortAndTokenOpenHarmony
}

const (
	// defaultTailLinkPort is the fixed loopback port that the TailLink
	// engine's LocalAPI server binds inside the VPN extension process.
	// Both sides default to it; TAILLINK_PORT can override either side.
	defaultTailLinkPort = 38288

	tailLinkPortEnv  = "TAILLINK_PORT"
	tailLinkTokenEnv = "TAILLINK_TOKEN"

	// tailLinkConfigRelPath is resolved against $HOME (set for the
	// Terminal-app user context) and holds the port/token shown by the
	// TailLink app's settings page.
	tailLinkConfigRelPath = ".config/taillink/cli.json"
)

type tailLinkCLIConfig struct {
	Port  int    `json:"port"`
	Token string `json:"token"`
}

var tailLinkCredsOnce struct {
	once  sync.Once
	port  int
	token string
	err   error
}

// localTCPPortAndTokenOpenHarmony returns the loopback TCP port and auth
// token of the TailLink engine's LocalAPI server.
//
// On OpenHarmony the engine runs inside the TailLink app's sandboxed VPN
// extension process and cannot expose a Unix socket in a shared location,
// so (like the macOS App Store variant) it serves LocalAPI over TCP on
// 127.0.0.1 and authenticates writes with a bearer token sent as the
// password of an HTTP Basic-Auth header.
//
// Credentials come from (highest precedence first):
//   - TAILLINK_PORT / TAILLINK_TOKEN environment variables
//   - $HOME/.config/taillink/cli.json ({"port":..., "token":...})
//   - the fixed default port, with no token
//
// A missing token is not an error: it limits the CLI to read-only
// endpoints (mirroring ipnserver's PermitRead/PermitWrite split).
func localTCPPortAndTokenOpenHarmony() (port int, token string, err error) {
	tailLinkCredsOnce.once.Do(func() {
		port, token, err := readTailLinkCreds()
		tailLinkCredsOnce.port, tailLinkCredsOnce.token, tailLinkCredsOnce.err = port, token, err
	})
	return tailLinkCredsOnce.port, tailLinkCredsOnce.token, tailLinkCredsOnce.err
}

func readTailLinkCreds() (port int, token string, err error) {
	if dir, err := os.UserHomeDir(); err == nil && dir != "" {
		if b, err := os.ReadFile(filepath.Join(dir, tailLinkConfigRelPath)); err == nil {
			var cfg tailLinkCLIConfig
			if json.Unmarshal(b, &cfg) == nil {
				port, token = cfg.Port, cfg.Token
			}
		}
	}

	if s := os.Getenv(tailLinkTokenEnv); s != "" {
		token = s
	}
	if s := os.Getenv(tailLinkPortEnv); s != "" {
		p, err := strconv.Atoi(s)
		if err != nil || p <= 0 || p > 65535 {
			return 0, "", fmt.Errorf("invalid %s value %q", tailLinkPortEnv, s)
		}
		port = p
	}

	if port == 0 {
		port = defaultTailLinkPort
	}
	return port, token, nil
}
