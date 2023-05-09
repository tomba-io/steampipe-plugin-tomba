package tomba

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableTombaVerifier(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tomba_verifier",
		Description: "Details about the deliverability of an email address.",
		List: &plugin.ListConfig{
			Hydrate: listVerifier,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name: "email"},
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
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the email address It takes 1 out of 6 possible values:valid,invalid,accept_all,webmail,disposable,unknown",
			},
			{
				Name:        "result",
				Type:        proto.ColumnType_STRING,
				Description: "status of the verification. It takes 1 out of 3 possible values: deliverable,undeliverable,risky",
			},
			{
				Name:        "score",
				Type:        proto.ColumnType_INT,
				Description: "deliverability score we give to the email address.",
			},
			// {
			// 	Name:        "data.sources",
			// 	Type:        proto.ColumnType_JSON,
			// 	Description: "The given email address somewhere on the web.",
			// },
		},
	}
}

func listVerifier(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tomba_verifier.listVerifier", "connection_error", err)
		return nil, err
	}

	email := d.EqualsQuals["email"].GetStringValue()
	plugin.Logger(ctx).Debug("tomba_verifier.listVerifier", "email", email)

	details, err := conn.EmailVerifier(email)
	plugin.Logger(ctx).Debug("tomba_verifier.listVerifier", "details", details)
	if err != nil {
		plugin.Logger(ctx).Error("tomba_verifier.listVerifier", "query_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, details.Data.Email)

	return nil, nil
}
