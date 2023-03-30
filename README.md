# Terraform Provider Portainer

[![Build](https://github.com/rogueai/terraform-provider-portainer/actions/workflows/go.yml/badge.svg)](https://github.com/rogueai/terraform-provider-portainer/actions/workflows/go.yml)

> Disclaimer: This provider is still a work in progress: do not use it in any production setting.
> Expect missing attributes, wrong updates, sudden changes in resource schema, etc.

## Status

Currently, the provider is able to manage only Swarm environments, with plain docker compose environments being next in line.

There are quite a few missing things: Environment Variables are not mapped, and the same applies for Resource Control attributes.

## Available Resources

These are the resources that are currently being worked on:

- Stack
- Secret

## Examples

Generated documentation is available, along with basic examples, under the `docs` folder.

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
