package main

import (
	"context"
	"flag"
	awx "github.com/ilijamt/terraform-provider-awx/internal/awx_21_8_0"
	"github.com/ilijamt/terraform-provider-awx/version"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/ilijamt/terraform-provider-awx/internal/provider"
)

func main() {
	var debug bool
	var err error

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	if err = providerserver.Serve(
		context.Background(),
		provider.NewFuncProvider(version.Version, nil, awx.Resources(), awx.DataSources()),
		providerserver.ServeOpts{
			Address: "registry.terraform.io/ilijamt/awx",
			Debug:   debug,
		},
	); err != nil {
		log.Fatal(err.Error())
	}
}
