package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsOrganizationsPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_organizations_policy",
		Description: "AWS Organizations Policy",
		Get: &plugin.GetConfig{
			Hydrate: kaytu.GetOrganizationsPolicy,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListOrganizationsPolicy,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name of the policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.PolicySummary.Name"),
			},
			{
				Name:        "id",
				Description: "The unique identifier (ID) of the policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.PolicySummary.Id"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.PolicySummary.Arn"),
			},
			{
				Name:        "type",
				Description: "The type of policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.PolicySummary.Type"),
			},
			{
				Name:        "aws_managed",
				Description: "A boolean value that indicates whether the specified policy is an Amazon Web Services managed policy. If true, then you can attach the policy to roots, OUs, or accounts, but you cannot edit it.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Policy.PolicySummary.AwsManaged"),
			},
			{
				Name:        "description",
				Description: "The description of the policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.PolicySummary.Description"),
			},
			{
				Name:        "content",
				Description: "The text content of the policy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy.Content"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.PolicySummary.Name"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy.PolicySummary.Arn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
