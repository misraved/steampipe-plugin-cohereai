package main

import (
	"github.com/Mr-Destructive/steampipe-plugin-cohereai/cohereai"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: cohere.Plugin})
}
