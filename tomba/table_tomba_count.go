package tomba

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTombaCount(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tomba_count",
		Description: "Details about total email addresses we have for one domain.",
		List: &plugin.ListConfig{
			Hydrate: listCount,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "domain", Require: plugin.Required},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "domain",
				Type:        proto.ColumnType_STRING,
				Hydrate:     domainString,
				Transform:   transform.FromValue(),
				Description: "The domain",
			},
			{
				Name:        "total",
				Type:        proto.ColumnType_INT,
				Description: "Total email on the domain.",
				Default:     0,
			},
			{
				Name:        "personal_emails",
				Type:        proto.ColumnType_INT,
				Description: "Total personal email on the domain.",
				Default:     0,
			},
			{
				Name:        "generic_emails",
				Type:        proto.ColumnType_INT,
				Description: "Total generic email on the domain.",
				Default:     0,
			},
		},
	}
}

func listCount(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tomba_count.listCount", "connection_error", err)
		return nil, err
	}

	domain := d.EqualsQuals["domain"].GetStringValue()
	plugin.Logger(ctx).Debug("tomba_count.listCount", "domain", domain)

	details, err := conn.Count(domain)
	plugin.Logger(ctx).Debug("tomba_count.listCount", "details", details)
	if err != nil {
		plugin.Logger(ctx).Error("tomba_count.listCount", "query_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, details.Data)

	return nil, nil
}
