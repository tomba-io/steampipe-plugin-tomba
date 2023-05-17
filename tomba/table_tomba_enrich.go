package tomba

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableTombaEnrich(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tomba_enrich",
		Description: "Details about person and company data based on an email, For example, you could retrieve a person’s name, location and social handles from an email.",
		List: &plugin.ListConfig{
			Hydrate: listEnrich,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "email"},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "email",
				Type:        proto.ColumnType_STRING,
				Description: "The email address.",
			},
			{
				Name:        "website_url",
				Type:        proto.ColumnType_STRING,
				Description: "The Company domain.",
			},
			{
				Name:        "company",
				Type:        proto.ColumnType_STRING,
				Description: "The company of person (if found)",
			},
			{
				Name:        "first_name",
				Type:        proto.ColumnType_STRING,
				Description: "The First name of person.",
			},
			{
				Name:        "last_name",
				Type:        proto.ColumnType_STRING,
				Description: "The Last name of person.",
			},
			{
				Name:        "gender",
				Type:        proto.ColumnType_STRING,
				Description: "The person gender.",
			},
			{
				Name:        "country",
				Type:        proto.ColumnType_STRING,
				Description: "The Two letter country code based on location.",
			},
			{
				Name:        "twitter",
				Type:        proto.ColumnType_STRING,
				Description: "The Twitter handle for the person (if found).",
			},
			{
				Name:        "linkedin",
				Type:        proto.ColumnType_STRING,
				Description: "The LinkedIn handle for the person (if found).",
			},
			{
				Name:        "sources",
				Type:        proto.ColumnType_JSON,
				Description: "The given email address somewhere on the web.",
			},
		},
	}
}

func listEnrich(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tomba_enrich.listEnrich", "connection_error", err)
		return nil, err
	}

	email := d.EqualsQuals["email"].GetStringValue()
	plugin.Logger(ctx).Debug("tomba_enrich.listEnrich", "email", email)

	details, err := conn.Enrichment(email)
	plugin.Logger(ctx).Debug("tomba_enrich.listEnrich", "details", details)
	if err != nil {
		plugin.Logger(ctx).Error("tomba_enrich.listEnrich", "query_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, details.Data)

	return nil, nil
}
