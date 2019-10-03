package main

import (
	"terraform-provider-junos/junos"

	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: junos.Provider,
	})
}