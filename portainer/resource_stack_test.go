package portainer

import (
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceStackSwarm(t *testing.T) {
	handlerFn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/api/stacks" {
			if r.Method == http.MethodPost {
				res := getFixture(t, "stacks_create")
				_, _ = w.Write(res)
				w.WriteHeader(http.StatusOK)
			}
		}
		if r.URL.Path == "/api/stacks/1" {
			if r.Method == http.MethodGet {
				res := getFixture(t, "stacks_read")
				_, _ = w.Write(res)
				w.WriteHeader(http.StatusOK)
			}
		}
		if r.URL.Path == "/api/stacks/1/file" {
			if r.Method == http.MethodGet {
				res := getFixture(t, "stacks_read_file")
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
			// Create and Read testing
			{
				Config: providerConfig + `
resource "portainer_stack" "test" {
  type              = 2
  endpoint_id       = 1
  name              = "myStack"
  file_content      = "version: 3\n services:\n web:\n image:nginx"
  swarm_id          = "jpofkc0i9uo9wtx1zesuk649w"
  namespace		    = ""
  env				= [{
	"name": "name"
	"value": "value"	
  }]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("portainer_stack.test", "name", "myStack"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("portainer_stack.test", "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "portainer_stack.test",
				ImportState:       true,
				ImportStateVerify: true,
				// The last_updated attribute does not exist in the Portainer
				// API, therefore there is no value for it during import.
				// ImportStateVerifyIgnore: []string{"last_updated"},
			},
			// Update and Read testing
			{
				Config: providerConfig + `
resource "portainer_stack" "test" {
  type              = 2
  endpoint_id       = 1
  name              = "myStack"
  file_content      = "version: 3\n services:\n web:\n image:nginx"
  swarm_id          = "jpofkc0i9uo9wtx1zesuk649w"
  namespace		    = ""
  env				= [{
	"name": "name"
	"value": "value"	
  }]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify first order item updated
					resource.TestCheckResourceAttr("portainer_stack.test", "id", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase

		},
	})
}
