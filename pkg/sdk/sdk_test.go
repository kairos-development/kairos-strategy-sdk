package sdk

import (
	"testing"

	"github.com/kairos-development/kairos-contracts/pkg/contracts"
)

func TestValidateManifest(t *testing.T) {
	valid := contracts.PluginManifest{
		Name:       "grid",
		Strategy:   "grid",
		Entrypoint: "on_tick",
		ABIVersion: contracts.ABIVersion{Major: 1, Minor: 0},
	}
	if err := ValidateManifest(valid); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tests := []struct {
		name     string
		manifest contracts.PluginManifest
	}{
		{name: "missing name", manifest: contracts.PluginManifest{Strategy: "grid", Entrypoint: "on_tick", ABIVersion: contracts.ABIVersion{Major: 1, Minor: 0}}},
		{name: "missing strategy", manifest: contracts.PluginManifest{Name: "grid", Entrypoint: "on_tick", ABIVersion: contracts.ABIVersion{Major: 1, Minor: 0}}},
		{name: "missing entrypoint", manifest: contracts.PluginManifest{Name: "grid", Strategy: "grid", ABIVersion: contracts.ABIVersion{Major: 1, Minor: 0}}},
		{name: "incompatible ABI", manifest: contracts.PluginManifest{Name: "grid", Strategy: "grid", Entrypoint: "on_tick", ABIVersion: contracts.ABIVersion{Major: 2, Minor: 0}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateManifest(tt.manifest); err == nil {
				t.Fatal("expected validation error")
			}
		})
	}
}
