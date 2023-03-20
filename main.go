package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	portainer "terraform-provider-portainer/portainer"
)

func main() {
	providerserver.Serve(context.Background(), portainer.New, providerserver.ServeOpts{
		// NOTE: This is not a typical Terraform Registry provider address,
		// such as registry.terraform.io/hashicorp/portainer. This specific
		// provider address is used in these tutorials in conjunction with a
		// specific Terraform CLI configuration for manual development testing
		// of this provider.
		Address: "rogueai.dev/rogueai/portainer",
	})
}
