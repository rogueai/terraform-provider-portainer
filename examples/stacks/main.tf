terraform {
  required_providers {
    portainer = {
      source = "rogueai/dev/portainer"
    }
  }
}

provider "portainer" {
  host    = "https://portainer.zugno.dev:8543/api"
  api_key = "ptr_3mTOfvl1UFKFz+DY033CZNcFCPQiMTm876Onl37Ik2g="
}

data "portainer_stacks" "list" {}

output "stacks" {
  value = data.portainer_stacks.list
}
