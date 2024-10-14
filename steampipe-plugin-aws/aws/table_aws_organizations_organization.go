package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsOrganizationsOrganization(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_organizations_organization",
		Description: "AWS Organizations Organization",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetOrganizationsOrganization,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListOrganizationsOrganization,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the organization.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Organization.Id")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the organization",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Organization.Arn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Organization.Id")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Organization.Arn").Transform(arnToAkas),
			},
		}),
	}
}
