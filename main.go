package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/xuyaming0800/terraform-provider-teststack/teststack"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: teststack.Provider,
	})
}
