package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-tomba/tomba"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: tomba.Plugin})
}
