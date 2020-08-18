# Standard Library Terraform Provider [![Build Status](https://travis-ci.com/invidian/terraform-provider-stdlib.svg?branch=master)](https://travis-ci.com/invidian/terraform-provider-stdlib) [![Maintainability](https://api.codeclimate.com/v1/badges/8a072bc32b8d61242133/maintainability)](https://codeclimate.com/github/invidian/terraform-provider-stdlib/maintainability) [![codecov](https://codecov.io/gh/invidian/terraform-provider-stdlib/branch/master/graph/badge.svg)](https://codecov.io/gh/invidian/terraform-provider-stdlib) [![Go Report Card](https://goreportcard.com/badge/github.com/invidian/terraform-provider-stdlib)](https://goreportcard.com/report/github.com/invidian/terraform-provider-stdlib)

This provider provides a collection of useful Terraform "functions", represented as a data sources, which are not part of the Terraform itself.

## Table of contents
* [User documentation](#user-documentation)
* [Building and testing](#building-and-testing)
* [Authors](#authors)

## User documentation

For user documentation, see [Terraform Registry](https://registry.terraform.io/providers/invidian/stdlib/latest/docs).

## Building

For testing builds, simply run `docker build .`, which will download all dependencies, run build, test and linter.

For local builds, run `make` which will build the binary, run unit tests and linter.

## Releasing

This project use `goreleaser` for releasing. To release new version, follow the following steps:

* Add a changelog for new release to CHANGELOG.md file.

* Tag new release on desired git, using example command:

  ```sh
  git tag -a v0.4.7 -s -m "Release v0.4.7"
  ```

* Push the tag to GitHub
  ```sh
  git push origin v0.4.7
  ```

* Run `goreleser` to create a GitHub Release:
  ```sh
  GITHUB_TOKEN=githubtoken GPG_FINGERPRINT=gpgfingerprint goreleaser release --release-notes <(go run github.com/rcmachado/changelog show 0.4.7)
  goreleaser
  ```

* Go to newly create [GitHub release](https://github.com/invidian/terraform-provider-stdlib/releases/tag/v0.4.7), verify that the changelog
  and artefacts looks correct and publish it.

## Authors

* **Mateusz Gozdek** - *Initial work* - [invidian](https://github.com/invidian)
