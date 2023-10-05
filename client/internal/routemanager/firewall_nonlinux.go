//go:build !linux && !ios
// +build !linux,!ios

package routemanager

import (
	"context"
	"fmt"
	"runtime"
)

// newFirewall returns a nil manager
func newFirewall(context.Context) (firewallManager, error) {
	return nil, fmt.Errorf("firewall not supported on %s", runtime.GOOS)
}
