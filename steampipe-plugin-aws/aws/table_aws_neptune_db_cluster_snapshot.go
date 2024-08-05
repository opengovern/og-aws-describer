package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	neptunev1 "github.com/aws/aws-sdk-go/service/neptune"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsNeptuneDBClusterSnapshot(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_neptune_db_cluster_snapshot",
		Description: "AWS Neptune DB Cluster Snapshot",
		Get: &plugin.GetConfig{
			Hydrate: kaytu.GetNeptuneDatabaseClusterSnapshot,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListNeptuneDatabaseClusterSnapshot,
		},
		GetMatrixItemFunc: SupportedRegionMatrix(neptunev1.EndpointsID),
		Columns: awsRegionalColumns([]*plugin.Column{
			{
				Name:        "db_cluster_snapshot_identifier",
				Description: "Specifies the identifier for a DB cluster snapshot. Must match the identifier of an existing snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.DBClusterSnapshotIdentifier"),
			},
			{
				Name:        "db_cluster_snapshot_arn",
				Description: "The Amazon Resource Name (ARN) for the DB cluster snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.DBClusterSnapshotArn"),
			},
			{
				Name:        "snapshot_type",
				Description: "Provides the type of the DB cluster snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.SnapshotType"),
			},
			{
				Name:        "status",
				Description: "Specifies the status of this DB cluster snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.Status"),
			},
			{
				Name:        "db_cluster_identifier",
				Description: "Specifies the DB cluster identifier of the DB cluster that this DB cluster snapshot was created from.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.DBClusterIdentifier"),
			},
			{
				Name:        "cluster_create_time",
				Description: "Specifies the time when the DB cluster was created, in Universal Coordinated Time (UTC).",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Snapshot.ClusterCreateTime"),
			},
			{
				Name:        "allocated_storage",
				Description: "Specifies the allocated storage size in gibibytes (GiB).",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Snapshot.AllocatedStorage"),
			},
			{
				Name:        "snapshot_create_time",
				Description: "Provides the time when the snapshot was taken, in Universal Coordinated Time (UTC).",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Snapshot.SnapshotCreateTime"),
			},
			{
				Name:        "percent_progress",
				Description: "Specifies the percentage of the estimated data that has been transferred.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Snapshot.PercentProgress"),
			},
			{
				Name:        "source_db_cluster_snapshot_arn",
				Description: "If the DB cluster snapshot was copied from a source DB cluster snapshot, the Amazon Resource Name (ARN) for the source DB cluster snapshot, otherwise, a null value.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.SourceDBClusterSnapshotArn"),
			},
			{
				Name:        "engine",
				Description: "Specifies the name of the database engine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.Engine"),
			},
			{
				Name:        "engine_version",
				Description: "Provides the version of the database engine for this DB cluster snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.EngineVersion"),
			},
			{
				Name:        "iam_database_authentication_enabled",
				Description: "True if mapping of Amazon Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Snapshot.IAMDatabaseAuthenticationEnabled"),
			},
			{
				Name:        "kms_key_id",
				Description: "If StorageEncrypted is true, the Amazon KMS key identifier for the encrypted DB cluster snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.KmsKeyId"),
			},
			{
				Name:        "license_model",
				Description: "Provides the license model information for this DB cluster snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.LicenseModel"),
			},
			{
				Name:        "master_username",
				Description: "Not supported by Neptune.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.MasterUsername"),
			},
			{
				Name:        "port",
				Description: "Specifies the port that the DB cluster was listening on at the time of the snapshot.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Snapshot.Port"),
			},
			{
				Name:        "storage_encrypted",
				Description: "Specifies whether the DB cluster snapshot is encrypted.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Snapshot.StorageEncrypted"),
			},
			{
				Name:        "vpc_id",
				Description: "Provides the VPC ID associated with the DB cluster snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.VpcId"),
			},
			{
				Name:        "availability_zones",
				Description: "Provides the list of EC2 Availability Zones that instances in the DB cluster snapshot can be restored in.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Snapshot.AvailabilityZones"),
			},
			{
				Name:        "db_cluster_snapshot_attributes",
				Description: "A list of DB cluster snapshot attribute names and values for a manual DB cluster snapshot.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Attributes"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.DBClusterSnapshotIdentifier"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Snapshot.DBClusterSnapshotArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
