// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build openharmony

package netns

import (
	"context"
	"net"
	"os"
	"testing"
)

func TestSetListenConfigInterfaceName(t *testing.T) {
	if os.Geteuid() != 0 {
		t.Skip("skipping; setting SO_BINDTODEVICE requires CAP_NET_RAW")
	}

	var lc net.ListenConfig
	if err := SetListenConfigInterfaceName(&lc, "lo"); err != nil {
		t.Fatal(err)
	}
	ln, err := lc.Listen(context.Background(), "tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	if _, ok := ln.(*net.TCPListener); !ok {
		t.Fatalf("got listener of type %T, want *net.TCPListener", ln)
	}
}
