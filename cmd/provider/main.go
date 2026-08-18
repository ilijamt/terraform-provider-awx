package main

import (
	"context"
	"flag"
	"log"

	"github.com/ilijamt/terraform-provider-awx/internal/awx"
	"github.com/ilijamt/terraform-provider-awx/version"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/ilijamt/terraform-provider-awx/internal/provider"
)

func main() {
	var debug bool
	var address string
	var err error

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.StringVar(&address, "address", "registry.terraform.io/ilijamt/awx", "provider address to serve under, e.g. registry.opentofu.org/ilijamt/awx to attach a debugger from tofu")
	flag.Parse()

	if err = providerserver.Serve(
		context.Background(),
		provider.NewFuncProvider(version.Version, nil, awx.Resources(), awx.DataSources()),
		providerserver.ServeOpts{
			Address: address,
			Debug:   debug,
		},
	); err != nil {
		log.Fatal(err.Error())
	}
}
