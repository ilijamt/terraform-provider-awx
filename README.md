AWX Terraform Provider
======================
[![Go Report Card](https://goreportcard.com/badge/github.com/ilijamt/terraform-provider-awx)](https://goreportcard.com/report/github.com/ilijamt/terraform-provider-awx)
[![Codecov](https://img.shields.io/codecov/c/gh/ilijamt/terraform-provider-awx)](https://app.codecov.io/gh/ilijamt/terraform-provider-awx)
[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/ilijamt/terraform-provider-awx)](go.mod)
[![GitHub](https://img.shields.io/github/license/ilijamt/terraform-provider-awx)](LICENSE)
[![Release](https://img.shields.io/github/release/ilijamt/terraform-provider-awx.svg)](https://github.com/ilijamt/terraform-provider-awx/releases/latest)

An autogenerated terraform provider based on the API specifications as provided by the `/api/v2/` endpoint.

AWX Versions
------------

Currently, built provider versions for AWX.

* 21.8.0
* 23.5.1 (only resources)
* 23.6.0 (only resources)

TODO:
-----

* Unit tests
* Integration tests

Download a new version of the API
---------------------------------

You need to spin up a version of AWX you want to download the API spec from.
Older version of AWX report incorrect API spec. So manual changes may be required to fix them.

```shell
export AWX_VERSION=23.6.0
mkdir -p resources/api/$AWX_VERSION/config
cat <<EOF > resources/api/$AWX_VERSION/config/default.json
{
  "api_version": "$AWX_VERSION"
}
EOF
node ./tools/config-merge.js $(pwd)/resources/config $(pwd)/resources/api/$AWX_VERSION
go run ./tools/generator/cmd/generator/main.go fetch-api-resources resources/api/$AWX_VERSION \
       --host $TOWER_HOST --password $TOWER_PASSWORD --username $TOWER_USERNAME --insecure-skip-verify
```

Build the version of the current API
-------------------------------------

```shell
make generate
```

If you want to build an API for a 23.6.0 just run
```shell
make generate VERSION=23.6.0
```