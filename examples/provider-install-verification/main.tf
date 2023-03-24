terraform {
  required_providers {
    portainer = {
      source = "rogueai/dev/portainer"
    }
  }
}

provider "portainer" {
  host = "<portainer-api-url>"
  api_key = "<portainer-api-key>"
}

data "portainer_stacks" "example" {}
