terraform {
  required_providers {
    portainer = {
      source = "github.com/rogueai/portainer"
    }
  }
}

provider "portainer" {
  host = "<portainer-api-url>"
  api_key = "<portainer-api-key>"
}

data "portainer_stacks" "list" {}

output "stacks" {
  value = data.portainer_stacks.list
}
