package main

import (
	"github.com/kaytu-io/kaytu-aws-describer/steampipe-plugin-aws/aws"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: aws.Plugin})
}
