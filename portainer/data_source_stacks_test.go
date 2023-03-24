package portainer

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceStacks(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + `data "portainer_stacks" "test" {}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify number of stacks returned
					resource.TestCheckResourceAttr("data.portainer_stacks.test", "stacks.#", "1"),
					//// Verify the first stack to ensure all attributes are set
					resource.TestCheckResourceAttr("data.portainer_stacks.test", "stacks.0.id", "1"),
					resource.TestCheckResourceAttr("data.portainer_stacks.test", "stacks.0.name", "myStack"),
					//// Verify placeholder id attribute
					resource.TestCheckResourceAttr("data.portainer_stacks.test", "id", "placeholder"),
				),
			},
		},
	})
}
