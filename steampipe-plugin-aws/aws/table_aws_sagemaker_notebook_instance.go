package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// // TABLE DEFINITION
func tableAwsSageMakerNotebookInstance(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_sagemaker_notebook_instance",
		Description: "AWS Sagemaker Notebook Instance",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationException", "NotFoundException", "RecordNotFound"}),
			},
			Hydrate: opengovernance.GetSageMakerNotebookInstance,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSageMakerNotebookInstance,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "default_code_repository", Require: plugin.Optional},
				{Name: "notebook_instance_lifecycle_config_name", Require: plugin.Optional},
				{Name: "notebook_instance_status", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the notebook instance.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.NotebookInstance.NotebookInstanceName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the notebook instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NotebookInstance.NotebookInstanceArn"),
			},
			{
				Name:        "creation_time",
				Description: "A timestamp that shows when the notebook instance was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.NotebookInstance.CreationTime")},
			{
				Name:        "default_code_repository",
				Description: "The Git repository associated with the notebook instance as its default code repository.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NotebookInstance.DefaultCodeRepository")},
			{
				Name:        "direct_internet_access",
				Description: "Describes whether Amazon SageMaker provides internet access to the notebook instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NotebookInstance.DirectInternetAccess")},
			{
				Name:        "failure_reason",
				Description: "If status is Failed, the reason it failed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NotebookInstance.FailureReason")},
			{
				Name:        "instance_type",
				Description: "The type of ML compute instance that the notebook instance is running on.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NotebookInstance.InstanceType")},
			{
				Name:        "kms_key_id",
				Description: "The AWS KMS key ID Amazon SageMaker uses to encrypt data when storing it on the ML storage volume attached to the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NotebookInstance.KmsKeyId")},
			{
				Name:        "last_modified_time",
				Description: "A timestamp that shows when the notebook instance was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.NotebookInstance.LastModifiedTime")},
			{
				Name:        "network_interface_id",
				Description: "The network interface IDs that Amazon SageMaker created at the time of creating the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NotebookInstance.NetworkInterfaceId")},
			{
				Name:        "notebook_instance_lifecycle_config_name",
				Description: "The name of a notebook instance lifecycle configuration associated with this notebook instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NotebookInstance.NotebookInstanceLifecycleConfigName")},
			{
				Name:        "notebook_instance_status",
				Description: "The status of the notebook instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NotebookInstance.NotebookInstanceStatus")},
			{
				Name:        "role_arn",
				Description: "The Amazon Resource Name (ARN) of the IAM role associated with the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NotebookInstance.RoleArn")},
			{
				Name:        "root_access",
				Description: "Whether root access is enabled or disabled for users of the notebook instance.Lifecycle configurations need root access to be able to set up a notebook instance",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NotebookInstance.RootAccess")},
			{
				Name:        "subnet_id",
				Description: "The ID of the VPC subnet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NotebookInstance.SubnetId")},
			{
				Name:        "url",
				Description: "The URL that you use to connect to the Jupyter notebook that is running in your notebook instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NotebookInstance.Url")},
			{
				Name:        "volume_size_in_gb",
				Description: "The size, in GB, of the ML storage volume attached to the notebook instance.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.NotebookInstance.VolumeSizeInGB")},
			{
				Name:        "accelerator_types",
				Description: "The list of the Elastic Inference (EI) instance types associated with this notebook instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NotebookInstance.AcceleratorTypes")},
			{
				Name:        "additional_code_repositories",
				Description: "An array of up to three Git repositories associated with the notebook instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NotebookInstance.AdditionalCodeRepositories")},
			{
				Name:        "security_groups",
				Description: "The IDs of the VPC security groups.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NotebookInstance.SecurityGroups")},
			{
				Name:        "tags_src",
				Description: "The list of tags for the notebook instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.NotebookInstance.NotebookInstanceName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(kaytuSageMakerNotebookInstanceTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NotebookInstance.NotebookInstanceArn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTION

func kaytuSageMakerNotebookInstanceTurbotTags(_ context.Context, d *transform.TransformData) (interface{},
	error) {
	tags := d.HydrateItem.(opengovernance.SageMakerNotebookInstance).Description.Tags

	if tags != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}

	return nil, nil
}
