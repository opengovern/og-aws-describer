package aws

import (
	"context"

	"github.com/kaytu-io/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSageMakerModel(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_sagemaker_model",
		Description: "AWS Sagemaker Model",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationException", "NotFoundException", "RecordNotFound"}),
			},
			Hydrate: opengovernance.GetSageMakerModel,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSageMakerModel,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "creation_time", Require: plugin.Optional, Operators: []string{">", ">=", "<", "<="}},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the model.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Model.ModelName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the model.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Model.ModelArn"),
			},
			{
				Name:        "creation_time",
				Description: "A timestamp that indicates when the model was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Model.CreationTime")},
			{
				Name:        "enable_network_isolation",
				Description: "If True, no inbound or outbound network calls can be made to or from the model container.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Model.EnableNetworkIsolation")},
			{
				Name:        "execution_role_arn",
				Description: "The Amazon Resource Name (ARN) of the IAM role that you specified for the model.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Model.ExecutionRoleArn")},
			{
				Name:        "containers",
				Description: "The containers in the inference pipeline.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Model.Containers")},
			{
				Name:        "inference_execution_config",
				Description: "Specifies details of how containers in a multi-container endpoint are called.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Model.InferenceExecutionConfig")},
			{
				Name:        "primary_container",
				Description: "The location of the primary inference code, associated artifacts, and custom environment map that the inference code uses when it is deployed in production.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Model.PrimaryContainer")},
			{
				Name:        "vpc_config",
				Description: "A VpcConfig object that specifies the VPC that this model has access to.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Model.VpcConfig")},
			{
				Name:        "tags_src",
				Description: "The list of tags for the model.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Model.ModelName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(sageMakerModelTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Model.ModelArn").Transform(arnToAkas),
			},
		}),
	}
}

func sageMakerModelTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.SageMakerModel).Description.Tags

	if tags != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}

	return nil, nil
}
