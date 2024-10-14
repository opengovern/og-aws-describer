package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsSsoAdminInstance(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ssoadmin_instance",
		Description: "AWS SSO Instance",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSSOAdminInstance,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The ARN of the SSO instance under which the operation will be executed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.InstanceArn"),
			},
			{
				Name:        "identity_store_id",
				Description: "The identifier of the identity store that is connected to the SSO instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.IdentityStoreId")},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Instance.InstanceArn")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Instance.InstanceArn").Transform(arnToAkas),
			},
		}),
	}
}
