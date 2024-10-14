package aws

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsFsxFileSystem(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_fsx_file_system",
		Description: "AWS FSx File System",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("file_system_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"FileSystemNotFound", "ValidationException"}),
			},
			Hydrate: opengovernance.GetFSXFileSystem,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListFSXFileSystem,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "file_system_id",
				Description: "The system-generated, unique 17-digit ID of the file system.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileSystem.FileSystemId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the EFS file system.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileSystem.ResourceARN"),
			},
			{
				Name:        "file_system_type",
				Description: "The type of Amazon FSx file system, which can be LUSTRE, WINDOWS, or ONTAP.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileSystem.FileSystemType")},
			{
				Name:        "lifecycle",
				Description: "The lifecycle status of the file system, following are the possible values AVAILABLE, CREATING, DELETING, FAILED, MISCONFIGURED, UPDATING.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileSystem.Lifecycle")},
			{
				Name:        "creation_time",
				Description: "The time that the file system was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.FileSystem.CreationTime")},
			{
				Name:        "dns_name",
				Description: "The DNS name for the file system.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileSystem.DNSName")},
			{
				Name:        "file_system_type_version",
				Description: "The version of your Amazon FSx for Lustre file system, either 2.10 or 2.12.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileSystem.FileSystemTypeVersion")},
			{
				Name:        "kms_key_id",
				Description: "The ID of the Key Management Service (KMS) key used to encrypt the file system's.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileSystem.KmsKeyId")},
			{
				Name:        "owner_id",
				Description: "The AWS account that created the file system.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileSystem.OwnerId")},
			{
				Name:        "storage_capacity",
				Description: "The storage capacity of the file system in gibibytes (GiB).",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.FileSystem.StorageCapacity")},
			{
				Name:        "storage_type",
				Description: "The storage type of the file system.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileSystem.StorageType")},
			{
				Name:        "vpc_id",
				Description: "The ID of the primary VPC for the file system.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileSystem.VpcId")},
			{
				Name:        "administrative_actions",
				Description: "A list of administrative actions for the file system that are in process or waiting to be processed.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FileSystem.AdministrativeActions")},
			{
				Name:        "failure_details",
				Description: "A structure providing details of any failures that occur when creating the file system has failed.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FileSystem.FailureDetails")},
			{
				Name:        "lustre_configuration",
				Description: "The configuration for the Amazon FSx for Lustre file system.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FileSystem.LustreConfiguration")},
			{
				Name:        "network_interface_ids",
				Description: "The IDs of the elastic network interface from which a specific file system is accessible.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FileSystem.NetworkInterfaceIds")},
			{
				Name:        "ontap_configuration",
				Description: "The configuration for this FSx for NetApp ONTAP file system.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FileSystem.OntapConfiguration")},
			{
				Name:        "open_zfs_configuration",
				Description: "The configuration for this FSx for NetApp ONTAP file system.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FileSystem.OpenZFSConfiguration"),
			},
			{
				Name:        "subnet_ids",
				Description: "Specifies the IDs of the subnets that the file system is accessible from.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FileSystem.SubnetIds")},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with Filesystem.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.FileSystem.Tags")},
			{
				Name:        "windows_configuration",
				Description: "The configuration for this Microsoft Windows file system.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FileSystem.WindowsConfiguration")},
			// Standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getFsxFileSystemTurbotTitle),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(fsxFileSystemTurbotData, "Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FileSystem.ResourceARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func fsxFileSystemTurbotData(_ context.Context, d *transform.TransformData) (interface{}, error) {
	fileSystemTag := d.HydrateItem.(opengovernance.FSXFileSystem).Description.FileSystem
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

func getFsxFileSystemTurbotTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	fileSystemTitle := d.HydrateItem.(opengovernance.FSXFileSystem).Description.FileSystem

	if fileSystemTitle.Tags != nil {
		for _, i := range fileSystemTitle.Tags {
			if *i.Key == "Name" && len(*i.Value) > 0 {
				return *i.Value, nil
			}
		}
	}

	return fileSystemTitle.FileSystemId, nil
}
