package portainer

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"os"
	"testing"
)

var (
	providerHost = "localhost:1234"
	// providerConfig is a shared configuration to combine with the actual
	// test configuration so the Portainer client is properly configured.
	// It is also possible to use the PORTAINER_ environment variables instead,
	// such as updating the Makefile and running the testing through that tool.
	providerConfig = fmt.Sprintf(`
provider "portainer" {
  host    = "http://%v/api"
  api_key = "ptr_abcdefgh123456"
}
`, providerHost)
)

var (
	// testAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"portainer": providerserver.NewProtocol6WithError(New()),
	}
	getFixture = func(t *testing.T, f string) []byte {
		file, err := os.ReadFile("../testdata/" + f + ".json")
		if err != nil {
			t.Errorf("Cannot load fixture: " + f + ", " + err.Error())
		}
		return file
	}
)
