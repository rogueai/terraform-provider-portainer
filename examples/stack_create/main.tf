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

resource "portainer_stack" "new" {
  type              = 1
  endpoint_id       = 2
  name              = "terraform_test"
  file_content      = file("docker-compose.yml")
  swarm_id          = "abcdefg123456"

}
