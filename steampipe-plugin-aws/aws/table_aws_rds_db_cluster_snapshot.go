package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRDSDBClusterSnapshot(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_rds_db_cluster_snapshot",
		Description: "AWS RDS DB Cluster Snapshot",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("db_cluster_snapshot_identifier"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"DBSnapshotNotFound", "DBClusterSnapshotNotFoundFault"}),
			},
			Hydrate: opengovernance.GetRDSDBClusterSnapshot,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRDSDBClusterSnapshot,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "db_cluster_identifier", Require: plugin.Optional},
				{Name: "db_cluster_snapshot_identifier", Require: plugin.Optional},
				{Name: "engine", Require: plugin.Optional},
				{Name: "type", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "db_cluster_snapshot_identifier",
				Description: "The friendly name to identify the DB Cluster Snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.DBClusterSnapshotIdentifier")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the DB Cluster Snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.DBClusterSnapshotArn"),
			},
			{
				Name:        "type",
				Description: "The type of the DB Cluster Snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.SnapshotType")},
			{
				Name:        "status",
				Description: "Specifies the status of this DB Cluster Snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.Status")},
			{
				Name:        "db_cluster_identifier",
				Description: "The friendly name to identify the DB Cluster, that the snapshot snapshot was created from.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.DBClusterIdentifier")},
			{
				Name:        "create_time",
				Description: "The time when the snapshot was taken.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DBClusterSnapshot.ClusterCreateTime")},
			{
				Name:        "allocated_storage",
				Description: "Specifies the allocated storage size in gibibytes (GiB).",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBClusterSnapshot.AllocatedStorage")},
			{
				Name:        "cluster_create_time",
				Description: "Specifies the time when the DB cluster was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DBClusterSnapshot.ClusterCreateTime")},
			{
				Name:        "engine",
				Description: "Specifies the name of the database engine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.Engine")},
			{
				Name:        "engine_version",
				Description: "Specifies the version of the database engine for this DB cluster snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.EngineVersion")},
			{
				Name:        "iam_database_authentication_enabled",
				Description: "Specifies whether mapping of AWS Identity and Access Management (IAM) accounts to database accounts is enabled, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBClusterSnapshot.IAMDatabaseAuthenticationEnabled")},
			{
				Name:        "kms_key_id",
				Description: "The AWS KMS key identifier for the AWS KMS customer master key (CMK).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.KmsKeyId")},
			{
				Name:        "license_model",
				Description: "Provides the license model information for this DB cluster snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.LicenseModel")},
			{
				Name:        "master_user_name",
				Description: "Provides the master username for the DB cluster snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.MasterUsername")},
			{
				Name:        "percent_progress",
				Description: "Specifies the percentage of the estimated data that has been transferred.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBClusterSnapshot.PercentProgress")},
			{
				Name:        "port",
				Description: "Specifies the port that the DB cluster was listening on at the time of the snapshot.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBClusterSnapshot.Port")},
			{
				Name:        "source_db_cluster_snapshot_arn",
				Description: "The Amazon Resource Name (ARN) for the source DB cluster snapshot, if the DB cluster snapshot was copied from a source DB cluster snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.SourceDBClusterSnapshotArn")},
			{
				Name:        "storage_encrypted",
				Description: "Specifies whether the DB cluster snapshot is encrypted, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBClusterSnapshot.StorageEncrypted")},
			{
				Name:        "vpc_id",
				Description: "Provides the VPC ID associated with the DB cluster snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.VpcId")},
			{
				Name:        "availability_zones",
				Description: "A list of Availability Zones (AZs) where instances in the DB cluster snapshot can be restored.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBClusterSnapshot.AvailabilityZones")},
			{
				Name:        "db_cluster_snapshot_attributes",
				Description: "A list of DB cluster snapshot attribute names and values for a manual DB cluster snapshot.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Attributes.DBClusterSnapshotAttributes")},
			{
				Name:        "tags_src",
				Description: "A list of tags attached to the DB Cluster Snapshot.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBClusterSnapshot.TagList"),
			},
			// Standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getRDSDBClusterSnapshotTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.DBClusterSnapshotIdentifier")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBClusterSnapshot.DBClusterSnapshotArn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func getRDSDBClusterSnapshotTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	dbClusterSnapshot := d.HydrateItem.(opengovernance.RDSDBClusterSnapshot).Description.DBClusterSnapshot

	if dbClusterSnapshot.TagList != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range dbClusterSnapshot.TagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}

//// UTILITY FUNCTIONS
