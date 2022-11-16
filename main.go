package main

import (
	"context"
	"flag"
	"github.com/ilijamt/terraform-provider-awx/config"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/ilijamt/terraform-provider-awx/internal/provider"
)

//go:generate go run ./tools/gen-awx-resources resources/api/21.5.0 internal/awx
//go:generate terraform fmt -recursive ./examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/ilijamt/awx",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(config.Version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
