# Upstream synchronization conflict report

> [!CAUTION]
> This is a temporary staging file. Do not merge this PR until the signed upstream release has been merged, every conflict has been resolved, this file has been removed, and the release manifest has been updated.

- Upstream release: [v1.104.1](https://github.com/tailscale/tailscale/releases/tag/v1.104.1)
- Upstream commit: `9a522a9786c97eb7910c01ccb7bd66557b04c910`
- Verified signing-key fingerprint: `SHA256:4X7QKZpmXLkns4z0XcXcUmifM/jruCmFyJnJIVmdUNo`
- Base branch: `main`
- Base commit: `a89dc45ba8903b25fc48827f7f7fbe47f539f1fa`
- Conflicting files: `252`
- Go/toolchain metadata changed: `true`

## Manual resolution

1. Check out this PR branch.
2. Fetch the verified tag and merge its peeled commit without auto-committing:

   ```sh
   git fetch https://github.com/tailscale/tailscale.git refs/tags/v1.104.1:refs/tags/v1.104.1
   git merge --no-ff --no-commit 9a522a9786c97eb7910c01ccb7bd66557b04c910
   ```

3. Resolve every conflict while preserving the fork-specific OpenHarmony adaptations. Do not apply a blanket `ours` or `theirs` resolution.
4. Remove this report as part of the signed-off upstream merge commit.
5. Update `.github/upstream-release.json`, run the required validation, and add the signed-off release-preparation commit.
6. Mark the PR ready only after the branch is complete and reviewed.

## Conflicting files

```text
.github/workflows/check-cherry-picks.yml
.github/workflows/test.yml
VERSION.txt
appc/appconnector.go
appc/conn25.go
appc/conn25_test.go
client/systray/systray.go
client/systray/systray_linux_test.go
client/web/auth.go
client/web/web_test.go
cmd/containerboot/main.go
cmd/containerboot/main_test.go
cmd/derper/depaware.txt
cmd/hello/helloserver/helloserver.go
cmd/k8s-operator/depaware.txt
cmd/k8s-operator/deploy/chart/templates/.gitignore
cmd/k8s-operator/deploy/crds/tailscale.com_peerrelays.yaml
cmd/k8s-operator/deploy/manifests/operator.yaml
cmd/k8s-operator/e2e/helpers.go
cmd/k8s-operator/e2e/ingress_test.go
cmd/k8s-operator/e2e/setup.go
cmd/k8s-operator/egress-pod-readiness.go
cmd/k8s-operator/egress-pod-readiness_test.go
cmd/k8s-operator/egress-services_test.go
cmd/k8s-operator/generate/main.go
cmd/k8s-operator/nodeport-service-ports.go
cmd/k8s-operator/nodeport-services-ports_test.go
cmd/k8s-operator/operator.go
cmd/k8s-operator/proxyclass.go
cmd/k8s-operator/proxyclass_test.go
cmd/k8s-operator/proxygroup.go
cmd/k8s-operator/proxygroup_test.go
cmd/k8s-operator/tsrecorder.go
cmd/k8s-operator/tsrecorder_test.go
cmd/proxy-to-grafana/proxy-to-grafana.go
cmd/stund/depaware.txt
cmd/tailscale/cli/debug.go
cmd/tailscale/cli/dns-query.go
cmd/tailscale/cli/dns-status.go
cmd/tailscale/cli/exitnode.go
cmd/tailscale/cli/funnel.go
cmd/tailscale/cli/jsonoutput/tailnet-lock-log.go
cmd/tailscale/cli/jsonoutput/tailnet-lock-status.go
cmd/tailscale/cli/routecheck.go
cmd/tailscale/cli/serve_legacy.go
cmd/tailscale/cli/serve_legacy_test.go
cmd/tailscale/cli/serve_v2.go
cmd/tailscale/cli/serve_v2_test.go
cmd/tailscale/cli/tailnet-lock.go
cmd/tailscale/cli/tailnet-lock_test.go
cmd/tailscale/cli/whois.go
cmd/tailscale/depaware.txt
cmd/tailscale/jsonoutput/example_format_test.go
cmd/tailscale/jsonoutput/example_responseenvelope_test.go
cmd/tailscale/jsonoutput/example_schemaversion_test.go
cmd/tailscale/jsonoutput/format_test.go
cmd/tailscale/jsonoutput/jsonoutput_test.go
cmd/tailscale/tsdnsjsonv0/example_dnsqueryresult_test.go
cmd/tailscale/tsdnsjsonv0/example_dnsstatusresult_test.go
cmd/tailscale/tsdnsjsonv0/status.go
cmd/tailscaled/depaware-min.txt
cmd/tailscaled/depaware-minbox.txt
cmd/tailscaled/depaware.txt
cmd/tailscaled/tailscaled.go
cmd/tailscaled/tailscaled_windows.go
cmd/testcontrol/testcontrol.go
cmd/tsconnect/serve.go
cmd/tsidp/depaware.txt
cmd/tsidp/tsidp.go
cmd/tsidp/tsidp_test.go
control/controlclient/auto_test.go
control/controlclient/direct.go
control/controlclient/direct_test.go
control/controlclient/map.go
control/controlclient/map_test.go
control/controlknobs/controlknobs.go
control/ts2021/client.go
control/ts2021/client_test.go
control/tsp/map.go
control/tsp/tsp.go
derp/client_test.go
derp/derphttp/derphttp_client.go
derp/derpserver/handler.go
drive/driveimpl/compositedav/stat_cache.go
drive/driveimpl/unicode_test.go
envknob/envknob.go
feature/ace/ace.go
feature/acme/cert.go
feature/acme/cert_test.go
feature/appconnectors/appconnectors.go
feature/c2n/c2n.go
feature/captiveportal/captiveportal.go
feature/conn25/api.go
feature/conn25/conn25.go
feature/conn25/conn25_test.go
feature/conn25/datapath_test.go
feature/debugportmapper/debugportmapper.go
feature/doctor/doctor.go
feature/featuretags/featuretags.go
feature/linkspeed/linkspeed_linux.go
feature/linuxdnsfight/linuxdnsfight.go
feature/portlist/portlist.go
feature/posture/posture.go
feature/relayserver/relayserver.go
feature/routecheck/routecheck.go
feature/routecheck/routertracker_test.go
feature/serviceclientprefs/serviceclientprefs.go
feature/ssh/ssh.go
feature/syspolicy/syspolicy.go
feature/taildrop/ext.go
feature/taildrop/peerapi.go
feature/taildrop/peerapi_test.go
feature/tailnetlock/tailnetlock.go
feature/tap/tap_linux.go
feature/wakeonlan/wakeonlan.go
flake.nix
flakehashes.json
go.mod
go.toolchain.next.rev
go.toolchain.rev
go.toolchain.version
ipn/backend.go
ipn/desktop/extension.go
ipn/ipn_clone.go
ipn/ipn_view.go
ipn/ipnext/ipnext.go
ipn/ipnlocal/bus.go
ipn/ipnlocal/bus_test.go
ipn/ipnlocal/dnsconfig_test.go
ipn/ipnlocal/drive.go
ipn/ipnlocal/drive_test.go
ipn/ipnlocal/drive_tomove.go
ipn/ipnlocal/local.go
ipn/ipnlocal/local_test.go
ipn/ipnlocal/netmapcache/netmapcache.go
ipn/ipnlocal/netmapcache/netmapcache_test.go
ipn/ipnlocal/node_backend.go
ipn/ipnlocal/node_backend_test.go
ipn/ipnlocal/peerapi.go
ipn/ipnlocal/peerapi_drive.go
ipn/ipnlocal/peerapi_test.go
ipn/ipnlocal/resolve.go
ipn/ipnlocal/serve.go
ipn/ipnlocal/serve_test.go
ipn/ipnlocal/tailnet-lock.go
ipn/ipnlocal/via.go
ipn/ipnlocal/via_test.go
ipn/ipnstate/ipnstate.go
ipn/localapi/debug.go
ipn/localapi/localapi_test.go
ipn/serve.go
ipn/serve_test.go
ipn/store/awsstore/store_aws.go
ipn/store/kubestore/store_kube.go
k8s-operator/api-proxy/proxy.go
k8s-operator/api-proxy/proxy_events_test.go
k8s-operator/api-proxy/proxy_test.go
k8s-operator/api.md
k8s-operator/apis/v1alpha1/types_peerrelay.go
k8s-operator/reconciler/dnsrecords/dnsrecords.go
k8s-operator/reconciler/dnsrecords/dnsrecords_test.go
k8s-operator/reconciler/nameserver/nameserver.go
k8s-operator/reconciler/nameserver/nameserver_test.go
k8s-operator/reconciler/peerrelay/authkey.go
k8s-operator/reconciler/peerrelay/peerrelay.go
k8s-operator/reconciler/peerrelay/service.go
k8s-operator/reconciler/peerrelay/statefulset.go
k8s-operator/reconciler/reconciler.go
k8s-operator/reconciler/reconciler_test.go
k8s-operator/reconciler/tailnet/tailnet.go
k8s-operator/reconciler/tailscaled/authkey.go
k8s-operator/sessionrecording/ws/frame_split_test.go
kube/state/state.go
net/dns/manager.go
net/dns/openresolv.go
net/dns/osconfig.go
net/dns/osconfig_test.go
net/dns/resolver/doh_test.go
net/dns/resolver/forwarder_test.go
net/dnscache/dnscache.go
net/netcheck/netcheck.go
net/netmon/netmon_test.go
net/packet/tsmp_test.go
net/routecheck/routecheck.go
net/socks5/socks5.go
net/stunserver/stunserver.go
net/tsdial/tsdial.go
net/tshttpproxy/tshttpproxy_windows.go
net/tstun/wrap.go
net/udprelay/server.go
safesocket/pipe_windows_test.go
shell.nix
ssh/tailssh/incubator.go
ssh/tailssh/incubator_plan9.go
ssh/tailssh/incubator_test.go
ssh/tailssh/localstate_test.go
ssh/tailssh/tailssh.go
ssh/tailssh/tailssh_integration_test.go
ssh/tailssh/tailssh_test.go
ssh/tailssh/user.go
tailcfg/tailcfg.go
tailcfg/tailcfg_clone.go
tailcfg/tailcfg_test.go
tailcfg/tailcfg_view.go
tka/sync.go
tka/tailchonk.go
tool/gocross/autoflags_test.go
tsnet/depaware.txt
tsnet/tailnetlock_test.go
tsnet/tsnet.go
tsnet/tsnet_test.go
tstest/integration/integration_test.go
tstest/integration/nat/nat_test.go
tstest/integration/netmapdelta_test.go
tstest/integration/service_windows.go
tstest/integration/tailscaled_deps_test_darwin.go
tstest/integration/tailscaled_deps_test_freebsd.go
tstest/integration/tailscaled_deps_test_linux.go
tstest/integration/tailscaled_deps_test_openbsd.go
tstest/integration/tailscaled_deps_test_windows.go
tstest/integration/testcontrol/testcontrol.go
tstest/largetailnet/delta_test.go
tstest/largetailnet/largetailnet.go
tstest/natlab/vmtest/connectivity_test.go
tstest/natlab/vmtest/dns_test.go
tstest/natlab/vmtest/drive_test.go
tstest/natlab/vmtest/ssh_test.go
tstest/natlab/vmtest/vmtest_test.go
tsweb/varz/varz.go
tsweb/varz/varz_test.go
types/netmap/netmap.go
util/deephash/deephash_test.go
util/linuxfw/nftables_runner_test.go
util/precompress/precompress.go
util/syspolicy/syspolicy_windows.go
util/zstdframe/options.go
wgengine/filter/filter_test.go
wgengine/filter/filtertype/filtertype.go
wgengine/filter/filtertype/filtertype_clone.go
wgengine/filter/match.go
wgengine/filter/tailcfg.go
wgengine/magicsock/endpoint.go
wgengine/magicsock/magicsock.go
wgengine/magicsock/magicsock_test.go
wgengine/magicsock/peermap.go
wgengine/magicsock/rebinding_conn.go
wgengine/netstack/netstack_test.go
wgengine/router/osrouter/router_freebsd.go
wgengine/router/osrouter/router_windows.go
wgengine/userspace.go
wgengine/wgcfg/config.go
wgengine/wgcfg/device.go
```
