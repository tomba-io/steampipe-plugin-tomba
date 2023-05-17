package tomba

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTombaAuthor(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tomba_author",
		Description: "Details about person and company data based on an URL.",
		List: &plugin.ListConfig{
			Hydrate: listAuthor,
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
				Description: "The URL of the article. For example, `https://clearbit.com/blog/company-name-to-domain-api`.",
			},
			{
				Name:        "email",
				Type:        proto.ColumnType_STRING,
				Description: "The email address.",
			},
			{
				Name:        "website_url",
				Type:        proto.ColumnType_STRING,
				Description: "The Domain of company’s website.",
			},
			{
				Name:        "company",
				Type:        proto.ColumnType_STRING,
				Description: "The Name of company (if found).",
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

func listAuthor(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tomba_author.listAuthor", "connection_error", err)
		return nil, err
	}

	url := d.EqualsQuals["url"].GetStringValue()
	plugin.Logger(ctx).Debug("tomba_author.listAuthor", "url", url)

	details, err := conn.AuthorFinder(url)
	plugin.Logger(ctx).Debug("tomba_author.listAuthor", "details", details)
	if err != nil {
		plugin.Logger(ctx).Error("tomba_author.listAuthor", "query_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, details.Data)

	return nil, nil
}
