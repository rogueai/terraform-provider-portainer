
# Creates an example stack.
resource "portainer_stack" "example" {
  type              = 1
  endpoint_id       = 2
  name              = "example"
  file_content      = file("docker-compose.yml")
  swarm_id          = "my-swarm-id"
}
