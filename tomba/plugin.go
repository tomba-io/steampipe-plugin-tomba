package tomba

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-tomba",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"tomba_account":  tableTombaAccount(ctx),
			"tomba_author":   tableTombaAuthor(ctx),
			"tomba_count":    tableTombaCount(ctx),
			"tomba_enrich":   tableTombaEnrich(ctx),
			"tomba_linkedin": tableTombaLinkedin(ctx),
			"tomba_verifier": tableTombaVerifier(ctx),
		},
	}
	return p
}
