package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsMediaStoreContainer(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_media_store_container",
		Description: "AWS Media Store Container",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameter", "ContainerNotFoundException", "ContainerInUseException"}),
			},
			Hydrate: kaytu.GetMediaStoreContainer,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListMediaStoreContainer,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ContainerInUseException"}),
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the container.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Container.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the container.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Container.ARN"),
			},
			{
				Name:        "status",
				Description: "The status of container creation or deletion. The status is one of the following: 'CREATING', 'ACTIVE', or 'DELETING'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Container.Status")},
			{
				Name:        "access_logging_enabled",
				Description: "The state of access logging on the container. This value is false by default, indicating that AWS Elemental MediaStore does not send access logs to Amazon CloudWatch Logs. When you enable access logging on the container, MediaStore changes this value to true, indicating that the service delivers access logs for objects stored in that container to CloudWatch Logs.",
				Type:        proto.ColumnType_BOOL,
				Default:     false,
				Transform:   transform.FromField("Description.Container.AccessLoggingEnabled")},
			{
				Name:        "creation_time",
				Description: "The Unix timestamp that the container was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Container.CreationTime")},
			{
				Name:        "endpoint",
				Description: "The DNS endpoint of the container.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Container.Endpoint")},
			{
				Name:        "policy",
				Description: "The contents of the access policy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy")},
			{
				Name:        "policy_std",
				Description: "Contains the contents of the access policy in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy").Transform(unescape).Transform(policyToCanonical),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with the container",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Container.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(containerTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Container.ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTION

func containerTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	item := d.HydrateItem.(kaytu.MediaStoreContainer).Description
	tags := item.Tags
	var turbotTagsMap map[string]string
	if tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}
	return turbotTagsMap, nil
}
