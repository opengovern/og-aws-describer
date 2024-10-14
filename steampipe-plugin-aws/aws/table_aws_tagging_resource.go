package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsTaggingResource(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_tagging_resource",
		Description: "AWS Tagging Resource",
		Get: &plugin.GetConfig{
			//Hydrate:    opengovernance.GetTaggingResources,
			KeyColumns: plugin.SingleColumn("arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterException"}),
			},
		},
		List: &plugin.ListConfig{
			//Hydrate: opengovernance.ListTaggingResources,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.TagMapping.ResourceARN").Transform(arnToTitle),
			},
			{
				Name:        "arn",
				Description: "The ARN of the resource.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.TagMapping.ResourceARN")},
			{
				Name:        "compliance_status",
				Description: "Whether a resource is compliant with the effective tag policy.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.TagMapping.ComplianceDetails.ComplianceStatus")},
			{
				Name:        "keys_with_noncompliant_values",
				Description: "These are keys defined in the effective policy that are on the resource with either incorrect case treatment or noncompliant values.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.TagMapping.ComplianceDetails.KeysWithNoncompliantValues")},
			{
				Name:        "noncompliant_keys",
				Description: "These tag keys on the resource are noncompliant with the effective tag policy.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.TagMapping.ComplianceDetails.NoncompliantKeys")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the parameter.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.TagMapping.Tags")},

			/// Steampipe standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.TagMapping.ResourceARN").Transform(arnToTitle),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.TagMapping.Tags").Transform(resourceTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.
					FromField("Description.TagMapping.ResourceARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

func resourceTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	tagList := d.Value.([]types.Tag)

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tagList != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
