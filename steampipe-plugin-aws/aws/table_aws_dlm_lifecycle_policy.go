package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsDLMLifecyclePolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_dlm_lifecycle_policy",
		Description: "AWS DLM Lifecycle Policy",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("policy_id"),
			Hydrate:    opengovernance.GetDLMLifecyclePolicy,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDLMLifecyclePolicy,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "policy_id",
				Description: "The identifier of the lifecycle policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LifecyclePolicy.PolicyId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LifecyclePolicy.PolicyArn"),
			},
			{
				Name:        "description",
				Description: "The description of the lifecycle policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LifecyclePolicy.Description")},
			{
				Name:        "date_created",
				Description: "The local date and time when the lifecycle policy was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.LifecyclePolicy.DateCreated")},
			{
				Name:        "date_modified",
				Description: "The local date and time when the lifecycle policy was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.LifecyclePolicy.DateModified")},
			{
				Name:        "execution_role_arn",
				Description: "The Amazon Resource Name (ARN) of the IAM role used to run the operations specified by the lifecycle policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LifecyclePolicy.ExecutionRoleArn")},
			{
				Name:        "policy_type",
				Description: "The type of policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LifecyclePolicy.PolicyDetails.PolicyType")},
			{
				Name:        "state",
				Description: "The activation state of the lifecycle policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LifecyclePolicy.State")},
			{
				Name:        "status_message",
				Description: "The description of the status.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LifecyclePolicy.StatusMessage")},
			{
				Name:        "policy_details",
				Description: "The configuration of the lifecycle policy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LifecyclePolicy.PolicyDetails")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LifecyclePolicy.Tags")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.LifecyclePolicy.PolicyId")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LifecyclePolicy.PolicyArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
