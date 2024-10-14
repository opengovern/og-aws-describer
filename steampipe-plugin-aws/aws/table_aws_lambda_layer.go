package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsLambdaLayer(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_lambda_layer",
		Description: "AWS Lambda Layer",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListLambdaLayer,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "layer_name",
				Description: "The name of the layer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Layer.LayerName")},
			{
				Name:        "layer_arn",
				Description: "The Amazon Resource Name (ARN) of the function layer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Layer.LayerArn")},
			{
				Name:        "created_date",
				Description: "The date that the version was created, in ISO 8601 format.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Layer.LatestMatchingVersion.CreatedDate")},
			{
				Name:        "description",
				Description: "The description of the version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Layer.LatestMatchingVersion.Description")},
			{
				Name:        "layer_version_arn",
				Description: "The ARN of the layer version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Layer.LatestMatchingVersion.LayerVersionArn")},
			{
				Name:        "license_info",
				Description: "The layer's open-source license.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Layer.LatestMatchingVersion.LicenseInfo")},
			{
				Name:        "version",
				Description: "The version number.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Layer.LatestMatchingVersion.Version")},
			{
				Name:        "compatible_architectures",
				Description: "A list of compatible instruction set architectures.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Layer.LatestMatchingVersion.CompatibleArchitectures")},
			{
				Name:        "compatible_runtimes",
				Description: "The layer's compatible runtimes.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Layer.LatestMatchingVersion.CompatibleRuntimes")},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Layer.LayerName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Layer.LayerArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
