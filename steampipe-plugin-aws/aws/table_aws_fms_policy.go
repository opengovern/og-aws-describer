package aws

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsFMSPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_fms_policy",
		Description: "AWS FMS Policy",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("policy_name"),
			//Hydrate:    opengovernance.GetFMSPolicy,
		},
		List: &plugin.ListConfig{
			//Hydrate: opengovernance.ListFMSPolicy,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.PolicyId")},
			{
				Name:        "policy_name",
				Description: "The name of the policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.PolicyName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the policy",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.PolicyArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.PolicyName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				//Transform:   transform.From(getFMSPolicyTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy.PolicyArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS
//
//func getFMSPolicyTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
//	tags := d.HydrateItem.(opengovernance.FMSPolicy).Description.Tags
//	return fmsV2TagsToMap(tags)
//}
