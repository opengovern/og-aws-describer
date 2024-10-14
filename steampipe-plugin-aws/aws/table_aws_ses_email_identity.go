package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsSESEmailIdentity(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ses_email_identity",
		Description: "AWS SES Email Identity",
		List: &plugin.ListConfig{
			Hydrate: ListSESEmailIdentities,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "identity",
				Description: "The email identity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Identity")},
			{
				Name:        "arn",
				Description: "The ARN of the AWS SES identity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "verification_status",
				Description: "The verification status of the identity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerificationAttributes.VerificationStatus")},
			{
				Name:        "verification_token",
				Description: "[DEPRECATED] This column has been deprecated and will be removed in a future release. The verification token for a domain identity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerificationAttributes.VerificationToken")},
			{
				Name:        "notification_attributes",
				Description: "Represents the notification attributes of an identity.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.NotificationAttributes")},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Identity")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

func ListSESEmailIdentities(ctx context.Context, d *plugin.QueryData, data *plugin.HydrateData) (interface{}, error) {
	d.EqualsQuals["identity_type"] = &proto.QualValue{
		Value: &proto.QualValue_StringValue{
			StringValue: string(types.IdentityTypeEmailAddress),
		},
	}
	return opengovernance.ListSESIdentity(ctx, d, data)
}
