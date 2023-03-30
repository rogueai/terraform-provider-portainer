
# Creates an example stack.
resource "portainer_secret" "example" {
  endpoint_id = 2
  name        = "mySecret"
  data        = base64encode("mySecretValue")
  labels      = {
    foo = "bar"
  }
}
