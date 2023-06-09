package portainer

import (
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceStacks(t *testing.T) {

	handlerFn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/api/stacks" {
			if r.Method == http.MethodGet {
				res := getFixture(t, "stacks_list")
				_, _ = w.Write(res)
				w.WriteHeader(http.StatusOK)
			}

		}
	}
	server := httptest.NewUnstartedServer(http.HandlerFunc(handlerFn))

	l, err := net.Listen("tcp", providerHost)
	if err != nil {
		log.Fatal(err)
	}
	_ = server.Listener.Close()
	server.Listener = l

	// Start the server.
	server.Start()
	// Stop the server on return from the function.
	defer server.Close()

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
