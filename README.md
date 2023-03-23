# Terraform Provider Portainer

## Examples
### Provider init
```terraform
terraform {
  required_providers {
    portainer = {
      source = "rogueai/dev/portainer"
    }
  }
}

provider "portainer" {
  host    = "<portainer_host>/api"
  api_key = "<portainer_api_key>"
}
```

### Stacks data source
```terraform

data "portainer_stacks" "list" {}

output "stacks" {
  value = data.portainer_stacks.list
}
```

### Stack resource creation

```terraform
resource "portainer_stack" "terraform_test" {
  type              = 1
  endpoint_id       = 2
  name              = "terraform_test"
  file_content      = file("docker-compose.yml")
  swarm_id          = "abcdefgh"

}
```
## Build provider

Run the following command to build the provider

```shell
$ go build -o terraform-provider-portainer
```

## Test sample configuration

First, build and install the provider.

```shell
$ make install
```

Then, navigate to the `examples` directory. 

```shell
$ cd examples
```

Run the following command to initialize the workspace and apply the sample configuration.

```shell
$ terraform init && terraform apply
```