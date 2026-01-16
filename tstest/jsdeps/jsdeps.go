// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Package jsdeps is a just a list of the packages we import in the
// JavaScript/WASM build, to let us test that our transitive closure of
// dependencies doesn't accidentally grow too large, since binary size
// is more of a concern.
package jsdeps

import (
	_ "bytes"
	_ "context"
	_ "encoding/hex"
	_ "encoding/json"
	_ "fmt"
	_ "log"
	_ "math/rand/v2"
	_ "net"
	_ "strings"
	_ "time"

	_ "golang.org/x/crypto/ssh"
	_ "github.com/Xinlong-Wu/tailscale-oh/control/controlclient"
	_ "github.com/Xinlong-Wu/tailscale-oh/ipn"
	_ "github.com/Xinlong-Wu/tailscale-oh/ipn/ipnserver"
	_ "github.com/Xinlong-Wu/tailscale-oh/net/netaddr"
	_ "github.com/Xinlong-Wu/tailscale-oh/net/netns"
	_ "github.com/Xinlong-Wu/tailscale-oh/net/tsdial"
	_ "github.com/Xinlong-Wu/tailscale-oh/safesocket"
	_ "github.com/Xinlong-Wu/tailscale-oh/tailcfg"
	_ "github.com/Xinlong-Wu/tailscale-oh/types/logger"
	_ "github.com/Xinlong-Wu/tailscale-oh/wgengine"
	_ "github.com/Xinlong-Wu/tailscale-oh/wgengine/netstack"
	_ "github.com/Xinlong-Wu/tailscale-oh/words"
)
