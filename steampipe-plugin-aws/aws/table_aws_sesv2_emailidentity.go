package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsSESv2EmailIdentity(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_sesv2_emailidentity",
		Description: "AWS SESv2 EmailIdentity",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetSESv2EmailIdentity,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSESv2EmailIdentity,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the emailidentity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Identity.IdentityName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the emailidentity",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ARN")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Identity.IdentityName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Identity.Tags"), // probably needs a transform function
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ARN").Transform(arnToAkas),
			},
		}),
	}
}
