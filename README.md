AWX Terraform Provider
======================
[![Go Report Card](https://goreportcard.com/badge/github.com/ilijamt/terraform-provider-awx)](https://goreportcard.com/report/github.com/ilijamt/terraform-provider-awx)
[![Codecov](https://img.shields.io/codecov/c/gh/ilijamt/terraform-provider-awx)](https://app.codecov.io/gh/ilijamt/terraform-provider-awx)
[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/ilijamt/terraform-provider-awx)](go.mod)
[![GitHub](https://img.shields.io/github/license/ilijamt/terraform-provider-awx)](LICENSE)
[![Release](https://img.shields.io/github/release/ilijamt/terraform-provider-awx.svg)](https://github.com/ilijamt/terraform-provider-awx/releases/latest)

An autogenerated terraform provider based on the API specifications as provided by the `/api/v2/` endpoint.

TODO:
-----
* Unit tests
* Integration tests

Download a new version of the API
---------------------------------

You need to spin up a version of AWX you want to download the API spec from. 
Older version of AWX report incorrect API spec. So manual changes may be required to fix them.

```shell
go run ./tools/generator/cmd/generator/main.go fetch-api-resources resources/config.json resources/api/21.8.0 \
       --host $TOWER_HOST --password $TOWER_PASSWORD --username $TOWER_USERNAME --insecure-skip-verify
```