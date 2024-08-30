package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsOrganizationsOrganizationalUnit(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_organizations_organizational_unit",
		Description: "AWS Organizations Organizational Unit",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    kaytu.GetOrganizationsOrganizationalUnit,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListOrganizationsOrganizationalUnit,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name of this OU.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Unit.Name"),
			},
			{
				Name:        "id",
				Description: "The unique identifier (ID) associated with this OU.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Unit.Id"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of this OU.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Unit.Arn"),
			},
			{
				Name:        "parent_id",
				Description: "The unique identifier (ID) of the root or OU whose child OUs you want to list.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ParentId"),
			},
			{
				Name:        "path",
				Description: "The OU path is a string representation that uniquely identifies the hierarchical location of an Organizational Unit within the AWS Organizations structure.",
				Type:        proto.ColumnType_LTREE,
				Transform:   transform.FromField("Description.Path"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Unit.Name"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Unit.Arn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
