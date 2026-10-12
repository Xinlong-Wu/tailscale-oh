// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package tailcfg aliases canonical Tailscale protocol types to this fork.
package tailcfg

import forktailcfg "github.com/Xinlong-Wu/tailscale-oh/tailcfg"

type (
	// CapGrant is the fork's capability grant type.
	CapGrant = forktailcfg.CapGrant
	// CapabilityVersion is the fork's protocol capability version.
	CapabilityVersion = forktailcfg.CapabilityVersion
	// DERPMap is the fork's DERP map type.
	DERPMap = forktailcfg.DERPMap
	// DERPRegionID is the fork's DERP region identifier.
	DERPRegionID = forktailcfg.DERPRegionID
	// DNSConfig is the fork's DNS configuration type.
	DNSConfig = forktailcfg.DNSConfig
	// DNSRecord is the fork's DNS record type.
	DNSRecord = forktailcfg.DNSRecord
	// FilterRule is the fork's packet-filter rule type.
	FilterRule = forktailcfg.FilterRule
	// Hostinfo is the fork's host information type.
	Hostinfo = forktailcfg.Hostinfo
	// HostinfoView is the fork's immutable host information view.
	HostinfoView = forktailcfg.HostinfoView
	// Login is the fork's login type.
	Login = forktailcfg.Login
	// LoginID is the fork's login identifier.
	LoginID = forktailcfg.LoginID
	// MapRequest is the fork's map request type.
	MapRequest = forktailcfg.MapRequest
	// MapResponse is the fork's map response type.
	MapResponse = forktailcfg.MapResponse
	// Node is the fork's node description type.
	Node = forktailcfg.Node
	// NodeCapMap is the fork's node capability map.
	NodeCapMap = forktailcfg.NodeCapMap
	// NodeID is the fork's node identifier.
	NodeID = forktailcfg.NodeID
	// OverTLSPublicKeyResponse is the fork's public-key response type.
	OverTLSPublicKeyResponse = forktailcfg.OverTLSPublicKeyResponse
	// PeerCapability is the fork's peer capability identifier.
	PeerCapability = forktailcfg.PeerCapability
	// PeerCapMap is the fork's peer capability map.
	PeerCapMap = forktailcfg.PeerCapMap
	// PeerChange is the fork's incremental peer update type.
	PeerChange = forktailcfg.PeerChange
	// PingRequest is the fork's ping request type.
	PingRequest = forktailcfg.PingRequest
	// RawMessage is the fork's raw JSON capability value.
	RawMessage = forktailcfg.RawMessage
	// RegisterRequest is the fork's registration request type.
	RegisterRequest = forktailcfg.RegisterRequest
	// RegisterResponse is the fork's registration response type.
	RegisterResponse = forktailcfg.RegisterResponse
	// SSHPolicy is the fork's SSH policy type.
	SSHPolicy = forktailcfg.SSHPolicy
	// SetDNSRequest is the fork's DNS update request type.
	SetDNSRequest = forktailcfg.SetDNSRequest
	// SetDNSResponse is the fork's DNS update response type.
	SetDNSResponse = forktailcfg.SetDNSResponse
	// StableNodeID is the fork's stable node identifier.
	StableNodeID = forktailcfg.StableNodeID
	// TKABootstrapResponse is the fork's Tailnet Lock bootstrap response.
	TKABootstrapResponse = forktailcfg.TKABootstrapResponse
	// TKAInfo is the fork's Tailnet Lock map information.
	TKAInfo = forktailcfg.TKAInfo
	// User is the fork's user type.
	User = forktailcfg.User
	// UserID is the fork's user identifier.
	UserID = forktailcfg.UserID
	// UserProfile is the fork's user profile type.
	UserProfile = forktailcfg.UserProfile
	// WebClientAuthResponse is the fork's web-client authorization response.
	WebClientAuthResponse = forktailcfg.WebClientAuthResponse
)

const (
	// CapabilitySSH is the canonical SSH node capability.
	CapabilitySSH = forktailcfg.CapabilitySSH
	// CurrentCapabilityVersion is the current protocol capability version.
	CurrentCapabilityVersion = forktailcfg.CurrentCapabilityVersion
)

// FilterAllowAll is the fork's allow-all packet filter.
var FilterAllowAll = forktailcfg.FilterAllowAll

// UnmarshalCapJSON decodes capability values using the fork implementation.
func UnmarshalCapJSON[T any](cm PeerCapMap, cap PeerCapability) ([]T, error) {
	return forktailcfg.UnmarshalCapJSON[T](cm, cap)
}
