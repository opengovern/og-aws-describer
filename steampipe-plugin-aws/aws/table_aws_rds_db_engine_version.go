package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/aws/aws-sdk-go-v2/service/rds/types"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRDSDBEngineVersion(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_rds_db_engine_version",
		Description: "AWS RDS DB Engine Version",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRDSDBEngineVersion,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "engine",
				Description: "The name of the database engine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EngineVersion.Engine"),
			},
			{
				Name:        "engine_version",
				Description: "The version number of the database engine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EngineVersion.EngineVersion"),
			},
			{
				Name:        "arn",
				Description: "The ARN of the custom engine version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EngineVersion.DBEngineVersionArn"),
			},
			{
				Name:        "status",
				Description: "The status of the DB engine version, either available or deprecated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EngineVersion.Status"),
			},
			{
				Name:        "create_time",
				Description: "The creation time of the DB engine version.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.EngineVersion.CreateTime"),
			},
			{
				Name:        "custom_db_engine_version_manifest",
				Description: "JSON string that lists the installation files and parameters that RDS Custom uses to create a custom engine version (CEV).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EngineVersion.CustomDBEngineVersionManifest"),
			},
			{
				Name:        "list_supported_character_sets",
				Description: "A value that indicates whether to list the supported character sets for each engine version.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.EngineVersion.CustomDBEngineVersionManifest"),
			},
			//{
			//	Name:        "engine_mode",
			//	Description: "Accepts DB engine modes.",
			//	Type:        proto.ColumnType_STRING,
			//	Transform:   transform.FromQual("engine_mode"),
			//},
			//{
			//	Name:        "list_supported_timezones",
			//	Description: "A value that indicates whether to list the supported time zones for each engine version.",
			//	Type:        proto.ColumnType_BOOL,
			//	Transform:   transform.FromQual("list_supported_timezones"),
			//	Default:     false,
			//},
			//{
			//	Name:        "default_only",
			//	Description: "A value that indicates whether only the default version of the specified engine or engine and major version combination is returned.",
			//	Type:        proto.ColumnType_BOOL,
			//	Transform:   transform.FromQual("default_only"),
			//	Default:     false,
			//},
			{
				Name:        "db_engine_description",
				Description: "The description of the database engine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EngineVersion.DBEngineDescription"),
			},
			{
				Name:        "db_engine_media_type",
				Description: "A value that indicates the source media provider of the AMI based on the usage operation. Applicable for RDS Custom for SQL Server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EngineVersion.DBEngineMediaType"),
			},
			{
				Name:        "db_engine_version_description",
				Description: "The description of the database engine version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EngineVersion.DBEngineVersionDescription"),
			},
			{
				Name:        "db_parameter_group_family",
				Description: "The name of the DB parameter group family for the database engine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EngineVersion.DBParameterGroupFamily"),
			},
			{
				Name:        "database_installation_files_s3_bucket_name",
				Description: "The name of the Amazon S3 bucket that contains your database installation files.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EngineVersion.DatabaseInstallationFilesS3BucketName"),
			},
			{
				Name:        "database_installation_files_s3_prefix",
				Description: "The Amazon S3 directory that contains the database installation files. If not specified, then no prefix is assumed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EngineVersion.DatabaseInstallationFilesS3Prefix"),
			},
			{
				Name:        "kms_key_id",
				Description: "The Amazon Web Services KMS key identifier for an encrypted CEV. This parameter is required for RDS Custom, but optional for Amazon RDS.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EngineVersion.KMSKeyId"),
			},
			{
				Name:        "major_engine_version",
				Description: "The major engine version of the CEV.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EngineVersion.MajorEngineVersion"),
			},
			{
				Name:        "supports_babelfish",
				Description: "A value that indicates whether the engine version supports Babelfish for Aurora PostgreSQL.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.EngineVersion.SupportsBabelfish"),
			},
			{
				Name:        "supports_certificate_rotation_without_restart",
				Description: "A value that indicates whether the engine version supports rotating the server certificate without rebooting the DB instance.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.EngineVersion.SupportsCertificateRotationWithoutRestart"),
			},
			{
				Name:        "supports_global_databases",
				Description: "A value that indicates whether you can use Aurora global databases with a specific DB engine version.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.EngineVersion.SupportsGlobalDatabases"),
			},
			{
				Name:        "supports_log_exports_to_cloudwatch_logs",
				Description: "A value that indicates whether the engine version supports exporting the log types specified by ExportableLogTypes to CloudWatch Logs.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.EngineVersion.SupportsLogExportsToCloudwatchLogs"),
			},
			{
				Name:        "supports_parallel_query",
				Description: "A value that indicates whether you can use Aurora parallel query with a specific DB engine version.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.EngineVersion.SupportsParallelQuery"),
			},
			{
				Name:        "supports_read_replica",
				Description: "Indicates whether the database engine version supports read replicas.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.EngineVersion.SupportsReadReplica"),
			},
			{
				Name:        "exportable_log_types",
				Description: "The types of logs that the database engine has available for export to CloudWatch Logs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EngineVersion.ExportableLogTypes"),
			},
			{
				Name:        "image",
				Description: "The EC2 image.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EngineVersion.Image"),
			},
			{
				Name:        "supported_feature_names",
				Description: "A list of features supported by the DB engine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EngineVersion.SupportedFeatureNames"),
			},
			{
				Name:        "supported_nchar_character_sets",
				Description: "A list of the character sets supported by the Oracle DB engine for the NcharCharacterSetName parameter of the CreateDBInstance operation.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EngineVersion.SupportedNcharCharacterSets"),
			},
			{
				Name:        "supported_timezones",
				Description: "A list of the time zones supported by this engine for the Timezone parameter of the CreateDBInstance action.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EngineVersion.SupportedTimezones"),
			},
			{
				Name:        "valid_upgrade_target",
				Description: "A list of engine versions that this database engine version can be upgraded to.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EngineVersion.ValidUpgradeTarget"),
			},
			{
				Name:        "tag_list",
				Description: "A list of tags.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EngineVersion.TagList"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EngineVersion.EngineVersion"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EngineVersion.TagList").Transform(getRDSDBEngineVersionTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EngineVersion.DBEngineVersionArn").Transform(getRDSDBEngineVersionAka),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getRDSDBEngineVersionTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	engineVersion := d.HydrateItem.(types.DBEngineVersion)

	if engineVersion.TagList != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range engineVersion.TagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}

func getRDSDBEngineVersionAka(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	engineVersion := d.HydrateItem.(types.DBEngineVersion)

	if engineVersion.DBEngineVersionArn == nil {
		return []string{}, nil
	} else {
		return transform.EnsureStringArray(ctx, d)
	}
}
