package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsSESDomainIdentity(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ses_domain_identity",
		Description: "AWS SES Domain Identity",
		List: &plugin.ListConfig{
			Hydrate: ListSESDomainIdentities,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "identity",
				Description: "The domain identity.",
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
				Description: "The verification token for a domain identity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerificationAttributes.VerificationToken"),
			},
			{
				//check the value of dkim attribute fields and the identity mail from domain attributes
				Name:        "dkim_attributes",
				Description: "The DKIM attributes for an email address or a domain.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DkimAttributes")},
			{
				Name:        "identity_mail_from_domain_attributes",
				Description: "The custom MAIL FROM attributes for a list of identities.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IdentityMail")},
			{
				Name:        "notification_attributes",
				Description: "Represents the notification attributes of an identity.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NotificationAttributes")},

			// Standard columns for all tables

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Identity")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

func ListSESDomainIdentities(ctx context.Context, d *plugin.QueryData, data *plugin.HydrateData) (interface{}, error) {
	d.EqualsQuals["identity_type"] = &proto.QualValue{
		Value: &proto.QualValue_StringValue{
			StringValue: string(types.IdentityTypeDomain),
		},
	}
	return opengovernance.ListSESIdentity(ctx, d, data)
}
