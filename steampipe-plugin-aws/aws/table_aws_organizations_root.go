package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsOrganizationsRoot(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_organizations_root",
		Description: "AWS Organizations Root",
		Get: &plugin.GetConfig{
			Hydrate: kaytu.GetOrganizationsRoot,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListOrganizationsRoot,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name of the root.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Root.Name"),
			},
			{
				Name:        "id",
				Description: "The unique identifier (ID) for the root.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Root.Id"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the root.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Root.Arn"),
			},
			{
				Name:        "policy_types",
				Description: "The types of policies that are currently enabled for the root and therefore can be attached to the root or to its OUs or accounts.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Root.PolicyTypes"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Root.Name"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Root.Arn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
