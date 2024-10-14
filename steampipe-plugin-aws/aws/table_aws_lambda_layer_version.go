package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsLambdaLayerVersion(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_lambda_layer_version",
		Description: "AWS Lambda Layer Version",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"layer_name", "version"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "InvalidParameter", "InvalidParameterValueException"}),
			},
			Hydrate: opengovernance.GetLambdaLayerVersion,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListLambdaLayerVersion,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "layer_name", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "layer_name",
				Description: "The name of the layer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LayerName")},
			{
				Name:        "layer_arn",
				Description: "The ARN of the layer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LayerVersion.LayerArn")},
			{
				Name:        "layer_version_arn",
				Description: "The ARN of the layer version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LayerVersion.LayerVersionArn")},
			{
				Name:        "created_date",
				Description: "The date that the version was created, in ISO 8601 format.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.LayerVersion.CreatedDate")},
			{
				Name:        "description",
				Description: "The description of the version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LayerVersion.Description")},
			{
				Name:        "license_info",
				Description: "The layer's open-source license.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LayerVersion.LicenseInfo")},
			{
				Name:        "revision_id",
				Description: "A unique identifier for the current revision of the policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.RevisionId")},
			{
				Name:        "version",
				Description: "The version number.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.LayerVersion.Version")},
			{
				Name:        "compatible_architectures",
				Description: "A list of compatible instruction set architectures.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LayerVersion.CompatibleArchitectures")},
			{
				Name:        "compatible_runtimes",
				Description: "The layer's compatible runtimes.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LayerVersion.CompatibleRuntimes")},
			{
				Name:        "content",
				Description: "Details about the layer version.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LayerVersion.Content")},
			{
				Name:        "policy",
				Description: "The policy document.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy")},
			{
				Name:        "policy_std",
				Description: "Contains the policy document in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy").Transform(unescape).Transform(policyToCanonical),
			},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LayerName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LayerVersion.LayerVersionArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
