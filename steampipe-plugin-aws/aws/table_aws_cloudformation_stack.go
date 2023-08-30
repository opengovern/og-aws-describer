package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCloudFormationStack(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudformation_stack",
		Description: "AWS CloudFormation Stack",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationError", "ResourceNotFoundException"}),
			},
			Hydrate: kaytu.GetCloudFormationStack,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCloudFormationStack,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Optional,
				},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "Unique identifier of the stack.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stack.StackId"),
			},
			{
				Name:        "name",
				Description: "The name associated with the stack.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Stack.StackName")},
			{
				Name:        "status",
				Description: "Current status of the stack.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Stack.StackStatus")},
			{
				Name:        "creation_time",
				Description: "The time at which the stack was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Stack.CreationTime")},
			{
				Name:        "disable_rollback",
				Description: "Boolean to enable or disable rollback on stack creation failures.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Stack.DisableRollback")},
			{
				Name:        "enable_termination_protection",
				Description: "Specifies whether termination protection is enabled for the stack.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Stack.EnableTerminationProtection")},
			{
				Name:        "last_updated_time",
				Description: "The time the stack was last updated. This field will only be returned if the stack has been updated at least once.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Stack.LastUpdatedTime")},
			{
				Name:        "parent_id",
				Description: "ID of the direct parent of this stack.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stack.ParentId")},
			{
				Name:        "role_arn",
				Description: "The Amazon Resource Name (ARN) of an AWS Identity and Access Management (IAM) role that is associated with the stack.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Stack.RoleARN")},
			{
				Name:        "root_id",
				Description: "ID of the top-level stack to which the nested stack ultimately belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stack.RootId")},
			{
				Name:        "description",
				Description: "A user-defined description associated with the stack.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stack.Description")},
			{
				Name:        "notification_arns",
				Description: "SNS topic ARNs to which stack related events are published.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Stack.NotificationARNs")},
			{
				Name:        "outputs",
				Description: "A list of output structures.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Stack.Outputs")},
			{
				Name:        "rollback_configuration",
				Description: "The rollback triggers for AWS CloudFormation to monitor during stack creation and updating operations, and for the specified monitoring period afterwards.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Stack.RollbackConfiguration")},
			{
				Name:        "capabilities",
				Description: "The capabilities allowed in the stack.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Stack.Capabilities")},
			{
				Name:        "stack_drift_status",
				Description: "Status of the stack's actual configuration compared to its expected template configuration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stack.DriftInformation.StackDriftStatus")},
			{
				Name:        "parameters",
				Description: "A list of Parameter structures.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Stack.Parameters")},
			{
				Name:        "template_body",
				Description: "Structure containing the template body.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StackTemplate.TemplateBody")},
			{
				Name:        "template_body_json",
				Description: "Structure containing the template body. Parsed into json object for better readability.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.StackTemplate.TemplateBody").Transform(UnmarshalYAMLorJSONNoUnescape),
			},
			{
				Name:        "resources",
				Description: "A list of Stack resource structures.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.StackResources")},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with stack.",
				Type:        proto.ColumnType_JSON,
				//Transform:   transform.From(cfnStackTagsToTurbotTags), //TODO-Saleh
				Transform: transform.FromField("Description.Stack.Tags")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Stack.Tags").Transform(stackTagListToTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Stack.StackName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Stack.StackId").Transform(arnToAkas),
			},
		}),
	}
}

func stackTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	if tagList, ok := d.Value.([]types.Tag); ok {
		// Mapping the resource tags inside turbotTags
		var turbotTagsMap map[string]string
		if tagList != nil {
			turbotTagsMap = map[string]string{}
			for _, i := range tagList {
				turbotTagsMap[*i.Key] = *i.Value
			}
		}
		return turbotTagsMap, nil
	} else {
		return nil, nil
	}
}
