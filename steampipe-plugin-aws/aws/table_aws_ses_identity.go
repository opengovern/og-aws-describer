package aws

import (
	"context"
	"fmt"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsSESIdentity(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ses_identity",
		Description: "AWS SES Identity",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("identity_name"),
			Hydrate:    kaytu.GetSESIdentity,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListSESIdentity,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "identity_name",
				Description: "The name of the identity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Identity.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the identity",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getSESIdentityArn)},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Identity.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getSESIdentityTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getSESIdentityArn).Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getSESIdentityTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.SESIdentity).Description.Tags
	return sesV2TagsToMap(tags)
}

func getSESIdentityArn(_ context.Context, d *transform.TransformData) (interface{}, error) {
	identity := d.HydrateItem.(kaytu.SESIdentity).Description.Identity
	metadata := d.HydrateItem.(kaytu.SESIdentity).Metadata

	arn := fmt.Sprintf("arn:%s:ses:%s:%s:identity/%s", metadata.Partition, metadata.Region, metadata.AccountID, identity)
	return arn, nil
}