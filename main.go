package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/invidian/terraform-provider-stdlib/stdlib"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: stdlib.Provider})
}
