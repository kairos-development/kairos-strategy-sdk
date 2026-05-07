package sdk

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kairos-development/kairos-contracts/pkg/contracts"
)

// SupportedABIMajor defines the major ABI version compatibility.
const (
	SupportedABIMajor = uint16(1)
)

// MinimumSupportedABI defines the minimal ABI version structure supported by the system.
var (
	MinimumSupportedABI = contracts.ABIVersion{Major: 1, Minor: 0}
)

// Context is the deterministic execution context injected into strategies.
type Context struct {
	NowUTC time.Time
	RNG    *rand.Rand
}

// Decision is the strategy output for a single tick.
type Decision struct {
	Signal *contracts.SignalEvent
	Intent *contracts.OrderIntentEvent
}

// Strategy is the public strategy SDK surface.
type Strategy interface {
	Manifest() contracts.PluginManifest
	OnTick(Context, contracts.MarketDataEvent) (Decision, error)
}

// ValidateManifest validates ABI compatibility and manifest completeness.
func ValidateManifest(manifest contracts.PluginManifest) error {
	if manifest.Name == "" {
		return fmt.Errorf("plugin manifest name is required")
	}
	if manifest.Strategy == "" {
		return fmt.Errorf("plugin manifest strategy is required")
	}
	if manifest.Entrypoint == "" {
		return fmt.Errorf("plugin manifest entrypoint is required")
	}
	if !manifest.ABIVersion.CompatibleWith(MinimumSupportedABI, SupportedABIMajor) {
		return fmt.Errorf("plugin ABI %s is incompatible with supported ABI %d.x", manifest.ABIVersion.String(), SupportedABIMajor)
	}
	return nil
}
