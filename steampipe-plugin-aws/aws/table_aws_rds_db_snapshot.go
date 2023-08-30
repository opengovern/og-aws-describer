package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRDSDBSnapshot(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_rds_db_snapshot",
		Description: "AWS RDS DB Snapshot",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("db_snapshot_identifier"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"DBSnapshotNotFound"}),
			},
			Hydrate: kaytu.GetRDSDBSnapshot,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListRDSDBSnapshot,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "db_instance_identifier", Require: plugin.Optional},
				{Name: "dbi_resource_id", Require: plugin.Optional},
				{Name: "engine", Require: plugin.Optional},
				{Name: "type", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "db_snapshot_identifier",
				Description: "The friendly name to identify the DB snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSnapshot.DBSnapshotIdentifier"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the DB snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSnapshot.DBSnapshotArn"),
			},
			{
				Name:        "type",
				Description: "Provides the type of the DB snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSnapshot.SnapshotType"),
			},
			{
				Name:        "status",
				Description: "Specifies the status of this DB snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSnapshot.Status"),
			},
			{
				Name:        "create_time",
				Description: "Specifies when the snapshot was taken.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DBSnapshot.SnapshotCreateTime"),
			},
			{
				Name:        "allocated_storage",
				Description: "Specifies the allocated storage size in gibibytes(GiB).",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBSnapshot.AllocatedStorage"),
			},
			{
				Name:        "availability_zone",
				Description: "Specifies the name of the Availability Zone the DB instance was located in, at the time of the DB snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSnapshot.AvailabilityZone"),
			},
			{
				Name:        "db_instance_identifier",
				Description: "Specifies the DB instance identifier of the DB instance this DB snapshot was created from.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSnapshot.DBInstanceIdentifier"),
			},
			{
				Name:        "dbi_resource_id",
				Description: "The identifier for the source DB instance, which can't be changed and which is unique to an AWS Region.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSnapshot.DbiResourceId"),
			},
			{
				Name:        "encrypted",
				Description: "Specifies whether the DB snapshot is encrypted, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBSnapshot.Encrypted"),
			},
			{
				Name:        "engine",
				Description: "Specifies the name of the database engine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSnapshot.Engine"),
			},
			{
				Name:        "engine_version",
				Description: "Specifies the version of the database engine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSnapshot.EngineVersion"),
			},
			{
				Name:        "iam_database_authentication_enabled",
				Description: "Specifies whether the mapping of AWS IAM accounts to database accounts is enabled, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBSnapshot.IAMDatabaseAuthenticationEnabled"),
			},
			{
				Name:        "instance_create_time",
				Description: "Specifies the time when the DB instance, from which the snapshot was taken, was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DBSnapshot.InstanceCreateTime"),
			},
			{
				Name:        "iops",
				Description: "Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBSnapshot.Iops"),
			},
			{
				Name:        "kms_key_id",
				Description: "Specifies the AWS KMS key identifier for the encrypted DB snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSnapshot.KmsKeyId"),
			},
			{
				Name:        "license_model",
				Description: "Specifies the License model information for the restored DB instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSnapshot.LicenseModel"),
			},
			{
				Name:        "master_user_name",
				Description: "Provides the master username for the DB snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSnapshot.MasterUsername"),
			},
			{
				Name:        "option_group_name",
				Description: "Provides the option group name for the DB snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSnapshot.OptionGroupName"),
			},
			{
				Name:        "percent_progress",
				Description: "The percentage of the estimated data that has been transferred.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBSnapshot.PercentProgress"),
			},
			{
				Name:        "port",
				Description: "Specifies the port that the database engine was listening on at the time of the snapshot.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBSnapshot.Port"),
			},
			{
				Name:        "source_db_snapshot_identifier",
				Description: "The DB snapshot ARN that the DB snapshot was copied from.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSnapshot.SourceDBSnapshotIdentifier"),
			},
			{
				Name:        "source_region",
				Description: "The AWS Region that the DB snapshot was created in or copied from.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSnapshot.SourceRegion"),
			},
			{
				Name:        "storage_type",
				Description: "Specifies the storage type associated with DB snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSnapshot.StorageType"),
			},
			{
				Name:        "tde_credential_arn",
				Description: "The ARN from the key store with which to associate the instance for TDE encryption.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSnapshot.TdeCredentialArn"),
			},
			{
				Name:        "timezone",
				Description: "The time zone of the DB snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSnapshot.Timezone"),
			},
			{
				Name:        "vpc_id",
				Description: "Provides the VPC ID associated with the DB snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSnapshot.VpcId"),
			},
			{
				Name:        "db_snapshot_attributes",
				Description: "A list of DB snapshot attribute names and values for a manual DB snapshot.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBSnapshotAttributes"),
			},
			{
				Name:        "processor_features",
				Description: "The number of CPU cores and the number of threads per core for the DB instance class of the DB instance when the DB snapshot was created.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBSnapshot.ProcessorFeatures"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags attached to the DB snapshot.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBSnapshot.TagList"),
			},

			// Standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getRDSDBSnapshotTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBSnapshot.DBSnapshotIdentifier"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBSnapshot.DBSnapshotArn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func getRDSDBSnapshotTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	dbSnapshot := d.HydrateItem.(kaytu.RDSDBSnapshot).Description.DBSnapshot

	if dbSnapshot.TagList != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range dbSnapshot.TagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}
