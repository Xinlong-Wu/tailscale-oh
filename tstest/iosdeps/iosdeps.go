// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Package iosdeps is a just a list of the packages we import on iOS, to let us
// test that our transitive closure of dependencies on iOS doesn't accidentally
// grow too large, as we've historically been memory constrained there.
package iosdeps

import (
	_ "bufio"
	_ "bytes"
	_ "context"
	_ "crypto/rand"
	_ "crypto/sha256"
	_ "encoding/json"
	_ "errors"
	_ "fmt"
	_ "io"
	_ "io/fs"
	_ "log"
	_ "math"
	_ "net"
	_ "net/http"
	_ "os"
	_ "os/signal"
	_ "path/filepath"
	_ "runtime"
	_ "runtime/debug"
	_ "strings"
	_ "sync"
	_ "sync/atomic"
	_ "syscall"
	_ "time"
	_ "unsafe"

	_ "github.com/tailscale/wireguard-go/device"
	_ "github.com/tailscale/wireguard-go/tun"
	_ "go4.org/mem"
	_ "golang.org/x/sys/unix"
	_ "github.com/Xinlong-Wu/tailscale-oh/hostinfo"
	_ "github.com/Xinlong-Wu/tailscale-oh/ipn"
	_ "github.com/Xinlong-Wu/tailscale-oh/ipn/ipnlocal"
	_ "github.com/Xinlong-Wu/tailscale-oh/ipn/localapi"
	_ "github.com/Xinlong-Wu/tailscale-oh/logtail"
	_ "github.com/Xinlong-Wu/tailscale-oh/logtail/filch"
	_ "github.com/Xinlong-Wu/tailscale-oh/net/dns"
	_ "github.com/Xinlong-Wu/tailscale-oh/net/netaddr"
	_ "github.com/Xinlong-Wu/tailscale-oh/net/tsdial"
	_ "github.com/Xinlong-Wu/tailscale-oh/net/tstun"
	_ "github.com/Xinlong-Wu/tailscale-oh/paths"
	_ "github.com/Xinlong-Wu/tailscale-oh/types/empty"
	_ "github.com/Xinlong-Wu/tailscale-oh/types/logger"
	_ "github.com/Xinlong-Wu/tailscale-oh/util/clientmetric"
	_ "github.com/Xinlong-Wu/tailscale-oh/util/dnsname"
	_ "github.com/Xinlong-Wu/tailscale-oh/version"
	_ "github.com/Xinlong-Wu/tailscale-oh/wgengine"
	_ "github.com/Xinlong-Wu/tailscale-oh/wgengine/router"
)
