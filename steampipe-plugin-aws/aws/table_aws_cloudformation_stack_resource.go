package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	cloudformationv1 "github.com/aws/aws-sdk-go/service/cloudformation"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCloudFormationStackResource(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudformation_stack_resource",
		Description: "AWS CloudFormation Stack Resource",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"stack_name", "logical_resource_id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationError", "ResourceNotFoundException"}),
			},
			Hydrate: kaytu.GetCloudFormationStackResource,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCloudFormationStackResource,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "stack_name",
					Require: plugin.Optional,
				},
			},
		},
		GetMatrixItemFunc: SupportedRegionMatrix(cloudformationv1.EndpointsID),
		Columns: awsRegionalColumns([]*plugin.Column{
			{
				Name:        "logical_resource_id",
				Description: "The logical name of the resource specified in the template.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StackResource.LogicalResourceId"),
			},
			{
				Name:        "stack_name",
				Description: "The name associated with the stack.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StackResource.StackName"),
			},
			{
				Name:        "stack_id",
				Description: "Unique identifier of the stack.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StackResource.StackId"),
			},
			{
				Name:        "last_updated_timestamp",
				Description: "Time the status was updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.StackResource.LastUpdatedTimestamp"),
			},
			{
				Name:        "resource_status",
				Description: "Current status of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StackResource.ResourceStatus").Transform(transform.ToString),
			},
			{
				Name:        "resource_type",
				Description: "Type of resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StackResource.ResourceType"),
			},
			{
				Name:        "description",
				Description: "User defined description associated with the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StackResource.Description"),
			},
			{
				Name:        "physical_resource_id",
				Description: "The name or unique identifier that corresponds to a physical instance ID of a resource supported by CloudFormation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StackResource.PhysicalResourceId"),
			},
			{
				Name:        "resource_status_reason",
				Description: "Success/failure message associated with the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StackResource.ResourceStatusReason"),
			},
			{
				Name:        "drift_information",
				Description: "Information about whether the resource's actual configuration differs, or has drifted, from its expected configuration, as defined in the stack template and any values specified as template parameters. For more information, see Detecting Unregulated Configuration Changes to Stacks and Resources.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.StackResource.DriftInformation"),
			},
			{
				Name:        "module_info",
				Description: "Contains information about the module from which the resource was created, if the resource was created from a module included in the stack template.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.StackResource.ModuleInfo"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StackResource.LogicalResourceId"),
			},
		}),
	}
}
