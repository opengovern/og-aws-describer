package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsCloudwatchLogResourcePolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudwatch_log_resource_policy",
		Description: "AWS CloudWatch Log Resource Policy",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCloudWatchLogResourcePolicy,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "policy_name",
				Description: "The name of the resource policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourcePolicy.PolicyName")},
			{
				Name:        "last_updated_time",
				Description: "Timestamp showing when this policy was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ResourcePolicy.LastUpdatedTime").Transform(convertTimestamp),
			},
			{
				Name:        "policy",
				Description: "The details of the policy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourcePolicy.PolicyDocument")},
			{
				Name:        "policy_std",
				Description: "Contains the policy document in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourcePolicy.PolicyDocument").Transform(unescape).Transform(policyToCanonical),
			},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourcePolicy.PolicyName")},
		}),
	}
}
