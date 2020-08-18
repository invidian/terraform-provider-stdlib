# Base64 Decode Hash Data Source

This data source takes Base64 encoded data as an argument and returns SHA1 and SHA256 hashes of given data
in attributes.

This data source can be used for example to calculate values for SSHFP DNS records based on PEM encoded OpenSSH Public Key.

## Example Usage

```hcl
data "stdlib_base64decodehash" "foo" {
  data = "tt8wbrUF4grdEoZP1kd3Ocqi7ncgtgq375sncoCm3ms="
}
```

## Argument Reference

* `data` - (Required) Base64 encoded data to hash.

## Attribute Reference

* `sha1` - SHA1 sum of given data.
* `sha256` - SHA256 sum of given data.
