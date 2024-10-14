package aws

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Not Mapped
// Don't need this one
func tableAwsDRSRecoverySnapshot(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_drs_recovery_snapshot",
		Description: "AWS DRS Recovery Snapshot",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "source_server_id", Require: plugin.Optional},
				{Name: "timestamp", Require: plugin.Optional, Operators: []string{">", ">=", "<", "<=", "="}},
			},
			IgnoreConfig: &plugin.IgnoreConfig{
				// UninitializedAccountException - This error comes up when default replication settings are not set for a particular region.
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"UninitializedAccountException", "BadRequestException"}),
			},
			//Hydrate: opengovernance.ListDRSRecoverySnapshot,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "snapshot_id",
				Description: "The ID of the snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RecoverySnapshot.SnapshotID")},
			{
				Name:        "source_server_id",
				Description: "The ID of the source server that the snapshot was taken for.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RecoverySnapshot.SourceServerID")},
			{
				Name:        "expected_timestamp",
				Description: "The timestamp of when we expect the snapshot to be taken.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.RecoverySnapshot.ExpectedTimestamp")},
			{
				Name:        "timestamp",
				Description: "The actual timestamp when the snapshot was taken.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.RecoverySnapshot.Timestamp")},
			{
				Name:        "ebs_snapshots",
				Description: "A list of EBS snapshots.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RecoverySnapshot.EbsSnapshots")},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RecoverySnapshot.SnapshotID")},
		}),
	}
}
