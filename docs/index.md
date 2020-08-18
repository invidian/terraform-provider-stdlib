# Standard Library Provider

The Standard Library Provider provides a collection of useful Terraform "functions", represented as a data sources, which are not part of the Terraform itself.

## Example Usage

```hcl
terraform {
  required_providers {
    sshcommand = {
      source  = "invidian/stdlib"
      version = "0.2.0"
    }
  }
}

data "stdlib_base64decodehash" "foo" {
  data = "tt8wbrUF4grdEoZP1kd3Ocqi7ncgtgq375sncoCm3ms="
}

output "sha1" {
  value = data.stdlib_base64decodehash.foo.sha1
}

output "sha256" {
  value = data.stdlib_base64decodehash.foo.sha256
}

```

## Argument Reference

This provider currently takes no arguments.
