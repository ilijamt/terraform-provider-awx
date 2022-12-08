package main

//go:generate terraform fmt -recursive ./examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	Version = "dev"
	//nolint:deadcode,unused,varcheck
	Commit = "SNAPSHOT"
)
