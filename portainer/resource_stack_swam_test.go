package portainer

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceStackSwarm(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
resource "portainer_stack" "test" {
  type              = 1
  endpoint_id       = 2
  name              = "terraform_test"
  file_content      = <<-EOT
    version: "3.7"
    
    services:
      alpine:
        image: alpine:latest
        tty: true
  EOT
  swarm_id          = "abcdefgh"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("portainer_stack.test", "name", "terraform_test"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("portainer_stack.test", "id"),
				),
			},
		},
	})
}
