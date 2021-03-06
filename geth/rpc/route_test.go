package rpc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// localMethods are methods that should be executed locally.
var localTestMethods = []string{"some_weirdo_method", "shh_newMessageFilter", "eth_accounts"}

func TestRouteWithUpstream(t *testing.T) {
	router := newRouter(true)

	for _, method := range remoteMethods {
		require.True(t, router.routeRemote(method), "method "+method+" should routed to remote")
	}

	for _, method := range localTestMethods {
		t.Run(method, func(t *testing.T) {
			require.False(t, router.routeRemote(method), "method "+method+" should routed to local")
		})
	}
}

func TestRouteWithoutUpstream(t *testing.T) {
	router := newRouter(false)

	for _, method := range remoteMethods {
		require.False(t, router.routeRemote(method), "method "+method+" should routed to locally without UpstreamEnabled")
	}

	for _, method := range localTestMethods {
		require.False(t, router.routeRemote(method), "method "+method+" should routed to local")
	}
}
