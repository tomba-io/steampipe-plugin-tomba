package tomba

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTombaLinkedin(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tomba_linkedin",
		Description: "Details about person and company data based on an Linkedin URL.",
		List: &plugin.ListConfig{
			Hydrate: listLinkedin,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "url"},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "url",
				Type:        proto.ColumnType_STRING,
				Hydrate:     urlString,
				Transform:   transform.FromValue(),
				Description: "The URL of the Linkedin. For example, `https://www.linkedin.com/in/alex-maccaw-ab592978`.",
			},
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

func listLinkedin(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tomba_linkedin.listLinkedin", "connection_error", err)
		return nil, err
	}

	url := d.EqualsQuals["url"].GetStringValue()
	plugin.Logger(ctx).Debug("tomba_linkedin.listLinkedin", "url", url)

	details, err := conn.LinkedinFinder(url)
	plugin.Logger(ctx).Debug("tomba_linkedin.listLinkedin", "details", details)
	if err != nil {
		plugin.Logger(ctx).Error("tomba_linkedin.listLinkedin", "query_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, details.Data)

	return nil, nil
}
