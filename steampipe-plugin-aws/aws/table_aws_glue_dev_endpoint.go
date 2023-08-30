package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsGlueDevEndpoint(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_glue_dev_endpoint",
		Description: "AWS Glue Dev Endpoint",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("endpoint_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"EntityNotFoundException"}),
			},
			Hydrate: kaytu.GetGlueDevEndpoint,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListGlueDevEndpoint,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "endpoint_name",
				Description: "The name of the DevEndpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DevEndpoint.EndpointName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the DevEndpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
			{
				Name:        "status",
				Description: "The current status of this DevEndpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DevEndpoint.Status")},
			{
				Name:        "availability_zone",
				Description: "The AWS Availability Zone where this DevEndpoint is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DevEndpoint.AvailabilityZone")},
			{
				Name:        "created_timestamp",
				Description: "The point in time at which this DevEndpoint was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DevEndpoint.CreatedTimestamp")},
			{
				Name:        "extra_jars_s3_path",
				Description: "The path to one or more Java .jar files in an S3 bucket that should be loaded in your DevEndpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DevEndpoint.ExtraJarsS3Path")},
			{
				Name:        "extra_python_libs_s3_path",
				Description: "The paths to one or more Python libraries in an Amazon S3 bucket that should be loaded in your DevEndpoint. Multiple values must be complete paths separated by a comma.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DevEndpoint.ExtraPythonLibsS3Path")},
			{
				Name:        "failure_reason",
				Description: "The reason for a current failure in this DevEndpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DevEndpoint.FailureReason")},
			{
				Name:        "glue_version",
				Description: "Glue version determines the versions of Apache Spark and Python that Glue supports.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DevEndpoint.GlueVersion")},
			{
				Name:        "last_modified_timestamp",
				Description: "The point in time at which this DevEndpoint was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DevEndpoint.LastModifiedTimestamp")},
			{
				Name:        "last_update_status",
				Description: "The status of the last update.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DevEndpoint.LastUpdateStatus")},
			{
				Name:        "number_of_nodes",
				Description: "The number of Glue Data Processing Units (DPUs) allocated to this DevEndpoint.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DevEndpoint.NumberOfNodes")},
			{
				Name:        "number_of_workers",
				Description: "The number of workers of a defined workerType that are allocated to the development endpoint.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DevEndpoint.NumberOfWorkers")},
			{
				Name:        "private_address",
				Description: "A private IP address to access the DevEndpoint within a VPC if the DevEndpoint is created within one.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DevEndpoint.PrivateAddress")},
			{
				Name:        "public_address",
				Description: "The public IP address used by this DevEndpoint. The PublicAddress field is present only when you create a non-virtual private cloud (VPC) DevEndpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DevEndpoint.PublicAddress")},
			{
				Name:        "public_key",
				Description: "The public key to be used by this DevEndpoint for authentication.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DevEndpoint.PublicKey")},
			{
				Name:        "role_arn",
				Description: "The Amazon Resource Name (ARN) of the IAM role used in this DevEndpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DevEndpoint.RoleArn")},
			{
				Name:        "security_configuration",
				Description: "The name of the SecurityConfiguration structure to be used with this DevEndpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DevEndpoint.SecurityConfiguration")},
			{
				Name:        "subnet_id",
				Description: "The subnet ID for this DevEndpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DevEndpoint.SubnetId")},
			{
				Name:        "vpc_id",
				Description: "The ID of the virtual private cloud (VPC) used by this DevEndpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DevEndpoint.VpcId")},
			{
				Name:        "worker_type",
				Description: "The type of predefined worker that is allocated to the development endpoint. Accepts a value of Standard, G.1X, or G.2X.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DevEndpoint.WorkerType")},
			{
				Name:        "yarn_endpoint_address",
				Description: "The YARN endpoint address used by this DevEndpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DevEndpoint.YarnEndpointAddress")},
			{
				Name:        "zeppelin_remote_spark_interpreter_port",
				Description: "The Apache Zeppelin port for the remote Apache Spark interpreter.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DevEndpoint.ZeppelinRemoteSparkInterpreterPort")},
			{
				Name:        "public_keys",
				Description: "A list of public keys to be used by the DevEndpoints for authentication.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DevEndpoint.PublicKeys")},
			{
				Name:        "security_group_ids",
				Description: "A list of security group identifiers used in this DevEndpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DevEndpoint.SecurityGroupIds")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DevEndpoint.EndpointName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
