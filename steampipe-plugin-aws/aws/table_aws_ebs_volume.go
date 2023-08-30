package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEBSVolume(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ebs_volume",
		Description: "AWS EBS Volume",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("volume_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidVolume.NotFound", "InvalidParameterValue"}),
			},
			Hydrate: kaytu.GetEC2Volume,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2Volume,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "availability_zone", Require: plugin.Optional},
				{Name: "encrypted", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "fast_restored", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "multi_attach_enabled", Require: plugin.Optional, Operators: []string{"=", "<>"}},
				{Name: "size", Require: plugin.Optional},
				{Name: "snapshot_id", Require: plugin.Optional},
				{Name: "state", Require: plugin.Optional},
				{Name: "volume_id", Require: plugin.Optional},
				{Name: "volume_type", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "volume_id",
				Description: "The ID of the volume.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Volume.VolumeId"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the volume.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN"),
			},
			{
				Name:        "volume_type",
				Description: "The volume type. This can be gp2 for General Purpose SSD, io1 or io2 for Provisioned IOPS SSD, st1 for Throughput Optimized HDD, sc1 for Cold HDD, or standard for Magnetic volumes.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Volume.VolumeType"),
			},
			{
				Name:        "state",
				Description: "The volume state.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Volume.State"),
			},
			{
				Name:        "create_time",
				Description: "The time stamp when volume creation was initiated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Volume.CreateTime"),
			},
			{
				Name:        "auto_enable_io",
				Description: "The state of autoEnableIO attribute.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Attributes.AutoEnableIO"),
			},
			{
				Name:        "availability_zone",
				Description: "The Availability Zone for the volume.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Volume.AvailabilityZone"),
			},
			{
				Name:        "encrypted",
				Description: "Indicates whether the volume is encrypted.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Volume.Encrypted"),
			},
			{
				Name:        "fast_restored",
				Description: "Indicates whether the volume was created using fast snapshot restore.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Volume.FastRestored"),
			},
			{
				Name:        "iops",
				Description: "The number of I/O operations per second (IOPS) that the volume supports.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Volume.Iops"),
			},
			{
				Name:        "kms_key_id",
				Description: "The Amazon Resource Name (ARN) of the AWS Key Management Service (AWS KMS) customer master key (CMK) that was used to protect the volume encryption key for the volume.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Volume.KmsKeyId"),
			},
			{
				Name:        "multi_attach_enabled",
				Description: "Indicates whether Amazon EBS Multi-Attach is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Volume.MultiAttachEnabled"),
			},
			{
				Name:        "outpost_arn",
				Description: "The Amazon Resource Name (ARN) of the Outpost.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Volume.OutpostArn"),
			},
			{
				Name:        "size",
				Description: "The size of the volume, in GiBs.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Volume.Size"),
			},
			{
				Name:        "snapshot_id",
				Description: "The snapshot from which the volume was created, if applicable.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Volume.SnapshotId"),
			},
			{
				Name:        "attachments",
				Description: "Information about the volume attachments.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Volume.Attachments"),
			},
			{
				Name:        "product_codes",
				Description: "A list of product codes.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Attributes.ProductCodes"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the volume.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Volume.Tags"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getEBSVolumeTitle),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEBSVolumeTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// LIST FUNCTION

// needed by other files

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func getEBSVolumeTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.EC2Volume).Description.Volume.Tags
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func getEBSVolumeTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	volume := d.HydrateItem.(kaytu.EC2Volume).Description.Volume

	if volume.Tags != nil {
		for _, i := range volume.Tags {
			if *i.Key == "Name" {
				return i.Value, nil
			}
		}
	}

	return volume.VolumeId, nil
}

//// UTILITY FUNCTION

// Build ebs volume list call input filter
