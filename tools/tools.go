//go:build tfplugindocs

// https://github.com/hashicorp/terraform-provider-hashicups/issues/60
package tools

import (
	// Ensure documentation generator is not removed from go.mod.
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
	// Ensure terraform-plugin-sdk is explicitly imported for testing requiring v2.29.0
	// @see: https://github.com/hashicorp/terraform-plugin-testing/issues/185
	_ "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)
