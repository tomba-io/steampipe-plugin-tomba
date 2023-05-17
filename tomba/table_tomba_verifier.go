package tomba

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTombaVerifier(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tomba_verifier",
		Description: "Details about the deliverability of an email address.",
		List: &plugin.ListConfig{
			Hydrate: listVerifier,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "email"},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "email",
				Type:        proto.ColumnType_STRING,
				Description: "The email address you want to verify.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the email address It takes 1 out of 6 possible values:valid,invalid,accept_all,webmail,disposable,unknown.",
			},
			{
				Name:        "result",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the verification. It takes 1 out of 3 possible values: deliverable,undeliverable,risky",
			},
			{
				Name:        "score",
				Type:        proto.ColumnType_INT,
				Description: "The deliverability score we give to the email address.",
			},
			{
				Name:        "accept_all",
				Type:        proto.ColumnType_BOOL,
				Description: "The SMTP server accepts all the email addresses. It means you can have have false positives on SMTP checks.",
			},
			{
				Name:        "block",
				Type:        proto.ColumnType_BOOL,
				Description: "The SMTP server prevented us to perform the SMTP check.",
			},
			{
				Name:        "registrar_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Whois.RegistrarName"),
				Description: "The domain name registrar is a business that handles the reservation of domain names as well as the assignment of IP addresses for those domain names",
			},
			{
				Name:        "created_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Whois.CreatedDate"),
				Description: "The creation date is when this domain was originally registered.",
			},
			{
				Name:        "referral_url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Whois.ReferralURL"),
				Description: "The RWhois (Referral Whois) is a directory services protocol which extends and enhances the Whois concept in a hierarchical and scalable fashion.",
			},
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
