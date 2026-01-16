// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Package licenses provides utilities for working with open source licenses.
package licenses

import "runtime"

// LicensesURL returns the absolute URL containing open source license information for the current platform.
func LicensesURL() string {
	switch runtime.GOOS {
	case "android":
		return "https://github.com/Xinlong-Wu/tailscale-oh/licenses/android"
	case "darwin", "ios":
		return "https://github.com/Xinlong-Wu/tailscale-oh/licenses/apple"
	case "windows":
		return "https://github.com/Xinlong-Wu/tailscale-oh/licenses/windows"
	default:
		return "https://github.com/Xinlong-Wu/tailscale-oh/licenses/tailscale"
	}
}
