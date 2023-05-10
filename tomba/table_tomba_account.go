package tomba

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
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
				Name:        "email",
				Type:        proto.ColumnType_STRING,
				Description: "The email address.",
			},
			{
				Name:        "secret_token",
				Type:        proto.ColumnType_STRING,
				Description: "The API secret.",
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
