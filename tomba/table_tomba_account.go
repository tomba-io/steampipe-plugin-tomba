package tomba

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTombaAccount(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tomba_account",
		Description: "Details about the current account",
		List: &plugin.ListConfig{
			Hydrate: listAccount,
		},
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "user_id",
				Type:        proto.ColumnType_INT,
				Description: "Your Account USER ID.",
			},
			{
				Name:        "first_name",
				Type:        proto.ColumnType_STRING,
				Description: "Your Account First name.",
			},
			{
				Name:        "last_name",
				Type:        proto.ColumnType_STRING,
				Description: "Your Account Last name.",
			},
			{
				Name:        "email",
				Type:        proto.ColumnType_STRING,
				Description: "Your Account email address.",
			},
			{
				Name:        "secret_token",
				Type:        proto.ColumnType_STRING,
				Description: "Your Account API secret.",
			},
			{
				Name:        "timezone",
				Type:        proto.ColumnType_STRING,
				Description: "Your Account timezone.",
			},
			{
				Name:        "country",
				Type:        proto.ColumnType_STRING,
				Description: "Your Account country, Your country is detected automatically.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date of your account's creation.",
			},
			{
				Name:        "pricing",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Pricing.Name"),
				Description: "The Subscription plan name.",
			},
		},
	}
}

func listAccount(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tomba_account.listAccount", "connection_error", err)
		return nil, err
	}

	details, err := conn.Account()
	plugin.Logger(ctx).Debug("tomba_account.listAccount", "details", details)
	if err != nil {
		plugin.Logger(ctx).Error("tomba_account.listAccount", "query_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, details.Data)

	return nil, nil
}
