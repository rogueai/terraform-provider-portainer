package portainer

import (
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceSecret(t *testing.T) {

	handlerFn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/api/endpoints/2/docker/secrets/create" {
			res := getFixture(t, "secret_create")
			_, _ = w.Write(res)
			w.WriteHeader(http.StatusCreated)
		}
		if r.URL.Path == "/api/endpoints/2/docker/secrets/mySecretId" {
			if r.Method == http.MethodGet {
				res := getFixture(t, "secret_read")
				_, _ = w.Write(res)
				w.WriteHeader(http.StatusOK)
			}
			if r.Method == http.MethodDelete {
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
resource "portainer_secret" "example" {
  endpoint_id = 2
  name        = "mySecret"
  data        = base64encode("mySecretValue")
  labels      = {
    foo = "bar"
  }
}
`,
			},
		},
	})
}
