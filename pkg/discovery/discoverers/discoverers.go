package discoverers

import (
	"context"

	invv1alpha1 "github.com/iptecharch/config-server/apis/inv/v1alpha1"
	"github.com/openconfig/gnmic/pkg/target"
)

var Discoverers = map[string]Initializer{}

type Initializer func() Discoverer

func Register(name string, initFn Initializer) {
	Discoverers[name] = initFn
}

// Discoverer discovers the target and returns discoveryInfo such as chassis type, SW version,
// SerialNumber, etc
type Discoverer interface {
	// Discover the target
	Discover(ctx context.Context, dr *invv1alpha1.DiscoveryRuleContext, t *target.Target) (*invv1alpha1.DiscoveryInfo, error)
	// GetName gets the provider name
	GetName() string
	// GetType gets the vendor device/NF type
	GetType() string
	// GetVendor gets the vendor name
	GetVendor() string
}
