//go:build tfplugindocs

// https://github.com/hashicorp/terraform-provider-hashicups/issues/60
package tools

import (
	// Ensure documentation generator is not removed from go.mod.
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
)
