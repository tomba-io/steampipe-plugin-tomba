package tomba

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type tombaConfig struct {
	Key    *string `cty:"key"`
	Secret *string `cty:"secret"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"key": {
		Type: schema.TypeString,
	},
	"secret": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &tombaConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) tombaConfig {
	if connection == nil || connection.Config == nil {
		return tombaConfig{}
	}
	config, _ := connection.Config.(tombaConfig)
	return config
}
