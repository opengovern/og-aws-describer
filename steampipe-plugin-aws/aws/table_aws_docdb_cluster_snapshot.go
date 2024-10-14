package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAwsDocDBClusterSnapshot(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_docdb_cluster_snapshot",
		Description: "AWS DocumentDB Cluster Snapshot",
		Get: &plugin.GetConfig{
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"DBSnapshotNotFound", "DBClusterSnapshotNotFoundFault"}),
			},
			Hydrate: opengovernance.GetDocDBClusterSnapshot,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDocDBClusterSnapshot,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "db_cluster_snapshot_identifier",
				Description: "The friendly name to identify the cluster snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.DBClusterIdentifier"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the cluster snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.DBClusterSnapshotArn"),
			},
			{
				Name:        "snapshot_type",
				Description: "The type of the cluster snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.SnapshotType"),
			},
			{
				Name:        "status",
				Description: "Specifies the status of this cluster snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.Status"),
			},
			//{
			//	Name:        "include_public",
			//	Description: "Set to true to include manual cluster snapshots that are public and can be copied or restored by any Amazon Web Services account, and otherwise false.",
			//	Type:        proto.ColumnType_BOOL,
			//	Transform:   transform.FromQual("include_public"),
			//	Default:     false,
			//},
			//{
			//	Name:        "include_shared",
			//	Description: "Set to true to include shared manual cluster snapshots from other Amazon Web Services accounts that this Amazon Web Services account has been given permission to copy or restore, and otherwise false.",
			//	Type:        proto.ColumnType_BOOL,
			//	Transform:   transform.FromQual("include_shared"),
			//	Default:     false,
			//},
			{
				Name:        "db_cluster_identifier",
				Description: "The friendly name to identify the cluster, that the snapshot snapshot was created from.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.DBClusterIdentifier"),
			},
			{
				Name:        "snapshot_create_time",
				Description: "The time when the snapshot was taken.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DBClusterSnapshot.SnapshotCreateTime"),
			},
			{
				Name:        "cluster_create_time",
				Description: "Specifies the time when the cluster was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DBClusterSnapshot.ClusterCreateTime"),
			},
			{
				Name:        "engine",
				Description: "Specifies the name of the database engine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.Engine"),
			},
			{
				Name:        "engine_version",
				Description: "Specifies the version of the database engine for this cluster snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.EngineVersion"),
			},
			{
				Name:        "kms_key_id",
				Description: "The AWS KMS key identifier for the AWS KMS customer master key (CMK).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.KmsKeyId"),
			},
			{
				Name:        "master_user_name",
				Description: "Provides the master username for the cluster snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.MasterUsername"),
			},
			{
				Name:        "percent_progress",
				Description: "Specifies the percentage of the estimated data that has been transferred.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBClusterSnapshot.PercentProgress"),
			},
			{
				Name:        "port",
				Description: "Specifies the port that the cluster was listening on at the time of the snapshot.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBClusterSnapshot.Port"),
			},
			{
				Name:        "source_db_cluster_snapshot_arn",
				Description: "The Amazon Resource Name (ARN) for the source cluster snapshot, if the cluster snapshot was copied from a source cluster snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.SourceDBClusterSnapshotArn"),
			},
			{
				Name:        "storage_encrypted",
				Description: "Specifies whether the cluster snapshot is encrypted, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBClusterSnapshot.StorageEncrypted"),
			},
			{
				Name:        "vpc_id",
				Description: "Provides the VPC ID associated with the cluster snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.VpcId"),
			},
			{
				Name:        "availability_zones",
				Description: "A list of Availability Zones (AZs) where instances in the cluster snapshot can be restored.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBClusterSnapshot.AvailabilityZones"),
			},
			{
				Name:        "db_cluster_snapshot_attributes",
				Description: "A list of DB cluster snapshot attribute names and values for a manual cluster snapshot.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Attributes.DBClusterSnapshotAttributes"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags attached to the cluster snapshot.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},

			// Standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(getDocDBClusterSnapshotTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBClusterSnapshot.DBClusterIdentifier"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBClusterSnapshot.DBClusterSnapshotArn").Transform(arnToAkas),
			},
		}),
	}
}

func getDocDBClusterSnapshotTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	tagList := d.Value.([]types.Tag)

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tagList != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
