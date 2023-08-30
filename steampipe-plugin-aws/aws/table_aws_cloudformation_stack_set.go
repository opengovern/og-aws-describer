package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCloudFormationStackSet(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudformation_stack_set",
		Description: "AWS CloudFormation StackSet",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("stack_set_name"),
			Hydrate:    kaytu.GetCloudFormationStackSet,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCloudFormationStackSet,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "status",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "stack_set_id",
				Description: "The ID of the stack set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StackSet.StackSetId"),
			},
			{
				Name:        "stack_set_name",
				Description: "The name of the stack set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StackSet.StackSetName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the stack set",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StackSet.StackSetARN"),
			},
			{
				Name:        "status",
				Description: "The status of the stack set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StackSet.Status"),
			},
			{
				Name:        "description",
				Description: "A description of the stack set that you specify when the stack set is created or updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StackSet.Description"),
			},
			{
				Name:        "drift_status",
				Description: "Status of the stack set's actual configuration compared to its expected template and parameter configuration. A stack set is considered to have drifted if one or more of its stack instances have drifted from their expected template and parameter configuration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StackSet.StackSetDriftDetectionDetails.DriftStatus"),
			},
			{
				Name:        "last_drift_check_timestamp",
				Description: "Most recent time when CloudFormation performed a drift detection operation on the stack set.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.StackSet.StackSetDriftDetectionDetails.LastDriftCheckTimestamp"),
			},
			{
				Name:        "permission_model",
				Description: "Describes how the IAM roles required for stack set operations are created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StackSet.PermissionModel")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getCloudFormationStackSetTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.StackSet.StackSetARN").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getCloudFormationStackSetTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.CloudFormationStackSet).Description.StackSet.Tags
	return cloudFormationV2TagsToMap(tags)
}
