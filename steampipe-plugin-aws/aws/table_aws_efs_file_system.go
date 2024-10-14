package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsElasticFileSystem(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_efs_file_system",
		Description: "AWS Elastic File System",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("file_system_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"FileSystemNotFound", "ValidationException"}),
			},
			Hydrate: opengovernance.GetEFSFileSystem,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEFSFileSystem,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "creation_token", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "Name of the file system provided by the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileSystem.Name"),
			},
			{
				Name:        "file_system_id",
				Description: "The ID of the file system, assigned by Amazon EFS.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileSystem.FileSystemId"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the EFS file system.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileSystem.FileSystemArn"),
			},
			{
				Name:        "owner_id",
				Description: "The AWS account that created the file system.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileSystem.OwnerId"),
			},
			{
				Name:        "creation_token",
				Description: "The opaque string specified in the request.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileSystem.CreationToken"),
			},
			{
				Name:        "creation_time",
				Description: "The time that the file system was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.FileSystem.CreationTime"),
			},
			{
				Name:        "automatic_backups",
				Description: "Automatic backups use a default backup plan with the AWS Backup recommended settings for automatic backups.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileSystem.Tags").Transform(automaticBackupsValue),
			},
			{
				Name:        "life_cycle_state",
				Description: "The lifecycle phase of the file system.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileSystem.LifeCycleState"),
			},
			{
				Name:        "number_of_mount_targets",
				Description: "The current number of mount targets that the file system has.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.FileSystem.NumberOfMountTargets"),
			},
			{
				Name:        "performance_mode",
				Description: "The performance mode of the file system.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileSystem.PerformanceMode"),
			},
			{
				Name:        "encrypted",
				Description: "A Boolean value that, if true, indicates that the file system is encrypted.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.FileSystem.Encrypted"),
			},
			{
				Name:        "kms_key_id",
				Description: "The ID of an AWS Key Management Service (AWS KMS) customer master key (CMK) that was used to protect the encrypted file system.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileSystem.KmsKeyId"),
			},
			{
				Name:        "throughput_mode",
				Description: "The throughput mode for a file system.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileSystem.ThroughputMode"),
			},
			{
				Name:        "provisioned_throughput_in_mibps",
				Description: "The throughput, measured in MiB/s, that you want to provision for a file system.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.FileSystem.ProvisionedThroughputInMibps"),
			},
			{
				Name:        "size_in_bytes",
				Description: "The latest known metered size (in bytes) of data stored in the file system.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FileSystem.SizeInBytes"),
			},
			{
				Name:        "policy",
				Description: "The JSON formatted FileSystemPolicy for the EFS file system.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy"),
			},
			{
				Name:        "policy_std",
				Description: "Contains the policy in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy").Transform(unescape).Transform(policyToCanonical),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with Filesystem.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FileSystem.Tags"),
			},

			// Standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getElasticFileSystemTurbotTitle),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FileSystem.Tags").Transform(elasticFileSystemTurbotData),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FileSystem.FileSystemArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func elasticFileSystemTurbotData(_ context.Context, d *transform.TransformData) (interface{}, error) {
	fileSystemTag := d.HydrateItem.(opengovernance.EFSFileSystem).Description.FileSystem
	if fileSystemTag.Tags == nil {
		return nil, nil
	}

	// Get the resource tags
	var turbotTagsMap map[string]string
	if fileSystemTag.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range fileSystemTag.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}
	return turbotTagsMap, nil
}

func getElasticFileSystemTurbotTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	fileSystemTitle := d.HydrateItem.(opengovernance.EFSFileSystem).Description.FileSystem

	if fileSystemTitle.Tags != nil {
		for _, i := range fileSystemTitle.Tags {
			if *i.Key == "Name" && len(*i.Value) > 0 {
				return *i.Value, nil
			}
		}
	}

	return fileSystemTitle.FileSystemId, nil
}

func automaticBackupsValue(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.EFSFileSystem).Description.FileSystem.Tags

	for _, i := range tags {
		if *i.Key == "aws:elasticfilesystem:default-backup" {
			return *i.Value, nil
		}
	}
	return "disabled", nil
}
