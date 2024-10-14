package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsEBSSnapshot(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ebs_snapshot",
		Description: "AWS EBS Snapshot",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("snapshot_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidSnapshot.NotFound", "InvalidSnapshotID.Malformed"}),
			},
			Hydrate: opengovernance.GetEC2VolumeSnapshot,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEC2VolumeSnapshot,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "description",
					Require: plugin.Optional,
				},
				{
					Name:    "encrypted",
					Require: plugin.Optional,
				},
				{
					Name:       "owner_alias",
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
				{
					Name:       "owner_id",
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
				{
					Name:       "snapshot_id",
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
				{
					Name:    "state",
					Require: plugin.Optional,
				},
				{
					Name:    "progress",
					Require: plugin.Optional,
				},
				{
					Name:    "volume_id",
					Require: plugin.Optional,
				},
				{
					Name:    "volume_size",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "snapshot_id",
				Description: "The ID of the snapshot. Each snapshot receives a unique identifier when it is created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.SnapshotId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN"),
			},
			{
				Name:        "state",
				Description: "The snapshot state.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.State")},
			{
				Name:        "volume_size",
				Description: "The size of the volume, in GiB.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Snapshot.VolumeSize")},
			{
				Name:        "volume_id",
				Description: "The ID of the volume that was used to create the snapshot. Snapshots created by the CopySnapshot action have an arbitrary volume ID that should not be used for any purpose.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.VolumeId")},
			{
				Name:        "encrypted",
				Description: "Indicates whether the snapshot is encrypted.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Snapshot.Encrypted")},
			{
				Name:        "start_time",
				Description: "The time stamp when the snapshot was initiated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Snapshot.StartTime")},
			{
				Name:        "description",
				Description: "The description for the snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.Description")},
			{
				Name:        "kms_key_id",
				Description: "The Amazon Resource Name (ARN) of the AWS Key Management Service (AWS KMS) customer master key (CMK) that was used to protect the volume encryption key for the parent volume.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.KmsKeyId")},
			{
				Name:        "data_encryption_key_id",
				Description: "The data encryption key identifier for the snapshot. This value is a unique identifier that corresponds to the data encryption key that was used to encrypt the original volume or snapshot copy. Because data encryption keys are inherited by volumes created from snapshots, and vice versa, if snapshots share the same data encryption key identifier, then they belong to the same volume/snapshot lineage.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.DataEncryptionKeyId")},
			{
				Name:        "progress",
				Description: "The progress of the snapshot, as a percentage.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.Progress")},
			{
				Name:        "state_message",
				Description: "Encrypted Amazon EBS snapshots are copied asynchronously. If a snapshot copy operation fails this field displays error state details to help you diagnose why the error occurred.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.StateMessage")},
			{
				Name:        "owner_alias",
				Description: "The AWS owner alias, from an Amazon-maintained list (amazon). This is not the user-configured AWS account alias set using the IAM console.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.OwnerAlias")},
			{
				Name:        "owner_id",
				Description: "The AWS account ID of the EBS snapshot owner.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.OwnerId")},
			{
				Name:        "create_volume_permissions",
				Description: "The users and groups that have the permissions for creating volumes from the snapshot.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.CreateVolumePermissions")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the snapshot.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Snapshot.Tags"),
			},

			/// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.SnapshotId"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(ec2SnapshotTurbotTags),
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

//// TRANSFORM FUNCTIONS

func ec2SnapshotTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.EC2VolumeSnapshot).Description.Snapshot.Tags

	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		if i.Key == nil || i.Value == nil {
			continue
		}
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}
