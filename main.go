package main

import (
	"github.com/tomba-io/steampipe-plugin-tomba/tomba"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: tomba.Plugin})
}
