package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsDmsEndpoint(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_dms_endpoint",
		Description: "AWS DMS Endpoint",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDMSEndpoint,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundFault", "InvalidParameterValueException"}),
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "endpoint_identifier",
				Description: "The database endpoint identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Endpoint.EndpointIdentifier"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) string that uniquely identifies the endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Endpoint.EndpointArn"),
			},
			{
				Name:        "certificate_arn",
				Description: "The Amazon Resource Name (ARN) used for SSL connection to the endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Endpoint.CertificateArn"),
			},
			{
				Name:        "database_name",
				Description: "The name of the database at the endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Endpoint.DatabaseName"),
			},
			{
				Name:        "endpoint_type",
				Description: "The type of endpoint. Valid values are source and target.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Endpoint.EndpointType"),
			},
			{
				Name:        "engine_display_name",
				Description: "The expanded name for the engine name. For example, if the EngineName parameter is 'aurora', this value would be 'Amazon Aurora MySQL'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Endpoint.EngineDisplayName"),
			},
			{
				Name:        "engine_name",
				Description: "The database engine name. Valid values, depending on the EndpointType, include 'mysql', 'oracle', 'postgres', 'mariadb', 'aurora', 'aurora-postgresql', 'redshift', 's3', 'db2', 'db2-zos', 'azuredb', 'sybase', 'dynamodb', 'mongodb', 'kinesis', 'kafka', 'elasticsearch', 'documentdb', 'sqlserver', 'neptune', and 'babelfish'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Endpoint.EngineName"),
			},
			{
				Name:        "external_id",
				Description: "Value returned by a call to CreateEndpoint that can be used for cross-account validation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Endpoint.ExternalId"),
			},
			{
				Name:        "external_table_definition",
				Description: "The external table definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Endpoint.ExternalTableDefinition"),
			},
			{
				Name:        "extra_connection_attributes",
				Description: "Additional connection attributes used to connect to the endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Endpoint.ExtraConnectionAttributes"),
			},
			{
				Name:        "kms_key_id",
				Description: "An KMS key identifier that is used to encrypt the connection parameters for the endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Endpoint.KmsKeyId"),
			},
			{
				Name:        "server_name",
				Description: "The name of the server at the endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Endpoint.ServerName"),
			},
			{
				Name:        "service_access_role_arn",
				Description: "The Amazon Resource Name (ARN) used by the service to access the IAM role.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Endpoint.ServiceAccessRoleArn"),
			},
			{
				Name:        "ssl_mode",
				Description: "The SSL mode used to connect to the endpoint. The default value is none.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Endpoint.SslMode"),
			},
			{
				Name:        "status",
				Description: "The status of the endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Endpoint.Status"),
			},
			{
				Name:        "username",
				Description: "The user name used to connect to the endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Endpoint.Username"),
			},
			{
				Name:        "port",
				Description: "The port value used to access the endpoint.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Endpoint.Port"),
			},
			// JSON columns
			{
				Name:        "dms_transfer_settings",
				Description: "The settings for the DMS Transfer type source.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Endpoint.DmsTransferSettings"),
			},
			{
				Name:        "doc_db_settings",
				Description: "Provides information that defines a DocumentDB endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Endpoint.DocDbSettings"),
			},
			{
				Name:        "dynamo_db_settings",
				Description: "The settings for the DynamoDB target endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Endpoint.DynamoDbSettings"),
			},
			{
				Name:        "elasticsearch_settings",
				Description: "The settings for the OpenSearch source endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Endpoint.ElasticsearchSettings"),
			},
			{
				Name:        "gcp_my_sql_settings",
				Description: "Settings in JSON format for the source GCP MySQL endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Endpoint.GcpMySQLSettings"),
			},
			{
				Name:        "ibm_db2_settings",
				Description: "The settings for the IBM Db2 LUW source endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Endpoint.IBMDb2Settings"),
			},
			{
				Name:        "kafka_settings",
				Description: "The settings for the Apache Kafka target endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Endpoint.KafkaSettings"),
			},
			{
				Name:        "kinesis_settings",
				Description: "The settings for the Amazon Kinesis target endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Endpoint.KinesisSettings"),
			},
			{
				Name:        "microsoft_sql_server_settings",
				Description: "The settings for the Microsoft SQL Server source and target endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Endpoint.MicrosoftSQLServerSettings"),
			},
			{
				Name:        "mongo_db_settings",
				Description: "The settings for the MongoDB source endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Endpoint.MongoDbSettings"),
			},
			{
				Name:        "my_sql_settings",
				Description: "The settings for the MySQL source and target endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Endpoint.MySQLSettings"),
			},
			{
				Name:        "neptune_settings",
				Description: "The settings for the Amazon Neptune target endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Endpoint.NeptuneSettings"),
			},
			{
				Name:        "oracle_settings",
				Description: "The settings for the Oracle source and target endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Endpoint.OracleSettings"),
			},
			{
				Name:        "postgre_sql_settings",
				Description: "The settings for the PostgreSQL source and target endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Endpoint.PostgreSQLSettings"),
			},
			{
				Name:        "redis_settings",
				Description: "The settings for the Redis target endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Endpoint.RedisSettings"),
			},
			{
				Name:        "redshift_settings",
				Description: "Settings for the Amazon Redshift endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Endpoint.RedshiftSettings"),
			},
			{
				Name:        "s3_settings",
				Description: "The settings for the S3 target endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Endpoint.S3Settings"),
			},
			{
				Name:        "sybase_settings",
				Description: "The settings for the SAP ASE source and target endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Endpoint.SybaseSettings"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags currently associated with the replication instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},

			// Steampipe Standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Endpoint.EndpointIdentifier"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(dmsEndpointTagListToTagsMap),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Endpoint.EndpointArn").Transform(arnToAkas),
			},
		}),
	}
}

func dmsEndpointTagListToTagsMap(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(*databasemigrationservice.ListTagsForResourceOutput)

	// Mapping the resource tags inside turbotTags
	if data.TagList != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range data.TagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}
