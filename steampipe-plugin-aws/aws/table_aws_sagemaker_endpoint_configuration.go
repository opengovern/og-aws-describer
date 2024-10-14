package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSageMakerEndpointConfiguration(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_sagemaker_endpoint_configuration",
		Description: "AWS Sagemaker Endpoint Configuration",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationException", "NotFoundException"}),
			},
			Hydrate: opengovernance.GetSageMakerEndpointConfiguration,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSageMakerEndpointConfiguration,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "creation_time", Require: plugin.Optional, Operators: []string{">", ">=", "<", "<="}},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the endpoint configuration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EndpointConfig.EndpointConfigName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the endpoint configuration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EndpointConfig.EndpointConfigArn"),
			},
			{
				Name:        "creation_time",
				Description: "A timestamp that shows when the endpoint configuration was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.EndpointConfig.CreationTime")},
			{
				Name:        "kms_key_id",
				Description: "AWS KMS key ID Amazon SageMaker uses to encrypt data when storing it on the ML storage volume attached to the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EndpointConfig.KmsKeyId")},
			{
				Name:        "data_capture_config",
				Description: "Specifies the parameters to capture input/output of Sagemaker models endpoints.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EndpointConfig.DataCaptureConfig")},
			{
				Name:        "production_variants",
				Description: "An array of ProductionVariant objects, one for each model that you want to host at this endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EndpointConfig.ProductionVariants")},
			{
				Name:        "tags_src",
				Description: "The list of tags for the endpoint configuration.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EndpointConfig.EndpointConfigName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(sageMakerEndpointConfigurationTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EndpointConfig.EndpointConfigArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func sageMakerEndpointConfigurationTurbotTags(_ context.Context, d *transform.TransformData) (interface{},
	error) {
	tags := d.HydrateItem.(opengovernance.SageMakerEndpointConfiguration).Description.Tags

	if tags != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}

	return nil, nil
}
