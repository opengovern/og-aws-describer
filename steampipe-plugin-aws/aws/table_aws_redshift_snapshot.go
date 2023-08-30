package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRedshiftSnapshot(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_redshift_snapshot",
		Description: "AWS Redshift Snapshot",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("snapshot_identifier"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ClusterSnapshotNotFound"}),
			},
			Hydrate: kaytu.GetRedshiftSnapshot,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListRedshiftSnapshot,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "cluster_identifier", Require: plugin.Optional},
				{Name: "owner_account", Require: plugin.Optional},
				{Name: "snapshot_type", Require: plugin.Optional},
				{Name: "snapshot_create_time", Require: plugin.Optional, Operators: []string{"="}},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "snapshot_identifier",
				Description: "The unique identifier of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.SnapshotIdentifier")},
			{
				Name:        "cluster_identifier",
				Description: "The identifier of the cluster for which the snapshot was taken.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.ClusterIdentifier")},
			{
				Name:        "snapshot_type",
				Description: "The snapshot type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.SnapshotType")},
			{
				Name:        "actual_incremental_backup_size_in_mega_bytes",
				Description: "The size of the incremental backup.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.Snapshot.ActualIncrementalBackupSizeInMegaBytes")},
			{
				Name:        "availability_zone",
				Description: "The Availability Zone in which the cluster was created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.AvailabilityZone")},
			{
				Name:        "backup_progress_in_mega-bytes",
				Description: "The number of megabytes that have been transferred to the snapshot backup.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.Snapshot.BackupProgressInMegaBytes")},
			{
				Name:        "cluster_create_time",
				Description: "The time (UTC) when the cluster was originally created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Snapshot.ClusterCreateTime")},
			{
				Name:        "cluster_version",
				Description: "The version ID of the Amazon Redshift engine that is running on the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.ClusterVersion")},
			{
				Name:        "current_backup_rate_in_mega_bytes_per_second",
				Description: "The number of megabytes per second being transferred to the snapshot backup.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.Snapshot.CurrentBackupRateInMegaBytesPerSecond")},
			{
				Name:        "db_name",
				Description: "The name of the database that was created when the cluster was created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.DBName")},
			{
				Name:        "elapsed_time_in_seconds",
				Description: "The amount of time an in-progress snapshot backup has been running, or the amount of time it took a completed backup to finish.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.ElapsedTimeInSeconds")},
			{
				Name:        "encrypted",
				Description: "If true, the data in the snapshot is encrypted at rest.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Snapshot.Encrypted")},
			{
				Name:        "encrypted_with_hsm",
				Description: "A boolean that indicates whether the snapshot data is encrypted using the HSM keys of the source cluster.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Snapshot.EncryptedWithHSM")},
			{
				Name:        "engine_full_version",
				Description: "The cluster version of the cluster used to create the snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.EngineFullVersion")},
			{
				Name:        "enhanced_vpc_routing",
				Description: "An option that specifies whether to create the cluster with enhanced VPC routing enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Snapshot.EnhancedVpcRouting")},
			{
				Name:        "estimated_seconds_to_completion",
				Description: "The estimate of the time remaining before the snapshot backup will complete.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.EstimatedSecondsToCompletion")},
			{
				Name:        "kms_key_id",
				Description: "The AWS KMS key ID of the encryption key that was used to encrypt data in the cluster from which the snapshot was taken.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.KmsKeyId")},
			{
				Name:        "maintenance_track_name",
				Description: "The name of the maintenance track for the snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.MaintenanceTrackName")},
			{
				Name:        "manual_snapshot_remaining_days",
				Description: "The number of days until a manual snapshot will pass its retention period.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Snapshot.ManualSnapshotRemainingDays")},
			{
				Name:        "manual_snapshot_retention_period",
				Description: "The number of days that a manual snapshot is retained.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Snapshot.ManualSnapshotRetentionPeriod")},
			{
				Name:        "master_username",
				Description: "The master user name for the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.MasterUsername")},
			{
				Name:        "node_type",
				Description: "The node type of the nodes in the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.NodeType")},
			{
				Name:        "number_of_nodes",
				Description: "The number of nodes in the cluster.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Snapshot.NumberOfNodes")},
			{
				Name:        "owner_account",
				Description: "The AWS customer account used to create or copy the snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.OwnerAccount")},
			{
				Name:        "port",
				Description: "The port that the cluster is listening on.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Snapshot.Port")},
			{
				Name:        "snapshot_create_time",
				Description: "The time (in UTC format) when Amazon Redshift began the snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.SnapshotCreateTime")},
			{
				Name:        "snapshot_retention_start_time",
				Description: "A timestamp representing the start of the retention period for the snapshot.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Snapshot.SnapshotRetentionStartTime")},
			{
				Name:        "source_region",
				Description: "The source region from which the snapshot was copied.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.SourceRegion")},
			{
				Name:        "status",
				Description: "The snapshot status.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.Status")},
			{
				Name:        "total_backup_size_in_mega_bytes",
				Description: "The size of the complete set of backup data that would be used to restore the cluster.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.Snapshot.TotalBackupSizeInMegaBytes")},
			{
				Name:        "vpc_id",
				Description: "The VPC identifier of the cluster if the snapshot is from a cluster in a VPC.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.VpcId")},
			{
				Name:        "accounts_with_restore_access",
				Description: "A list of the AWS customer accounts authorized to restore the snapshot.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Snapshot.AccountsWithRestoreAccess")},
			{
				Name:        "restorable_node_types",
				Description: "The list of node types that this cluster snapshot is able to restore into.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Snapshot.RestorableNodeTypes")},
			{
				Name:        "tags_src",
				Description: "The list of tags for the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Snapshot.Tags")},

			// Standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.SnapshotIdentifier")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(redshiftSnapshotTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func redshiftSnapshotTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	snapshot := d.HydrateItem.(kaytu.RedshiftSnapshot).Description.Snapshot

	// Get the resource tags
	if snapshot.Tags != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range snapshot.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}
