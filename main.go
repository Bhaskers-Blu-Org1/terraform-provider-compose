package main

import (
	"github.com/IBM/terraform-provider-compose/compose"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: compose.Provider})
}
