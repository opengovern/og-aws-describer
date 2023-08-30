package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsKeyspacesKeyspace(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_keyspaces_keyspace",
		Description: "AWS Keyspaces Keyspace",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("keyspace_name"),
			Hydrate:    kaytu.GetKeyspacesKeyspace,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListKeyspacesKeyspace,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "keyspace_name",
				Description: "The name of the keyspace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Keyspace.KeyspaceName")},
			{
				Name:        "resource_arn",
				Description: "The Amazon Resource Name (ARN) of the keyspace",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Keyspace.ResourceArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Keyspace.KeyspaceName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getKeyspacesKeyspaceTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Keyspace.ResourceArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getKeyspacesKeyspaceTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.KeyspacesKeyspace).Description.Tags
	return keyspacesV2TagsToMap(tags)
}
