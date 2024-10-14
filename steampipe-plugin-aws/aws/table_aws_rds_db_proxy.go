package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRDSDBProxy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_rds_db_proxy",
		Description: "AWS RDS DB Proxy",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("db_proxy_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"DBProxyNotFoundFault"}),
			},
			Hydrate: opengovernance.GetRDSDBProxy,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRDSDBProxy,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidAction"}),
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "db_proxy_name",
				Description: "The identifier for the proxy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBProxy.DBProxyName")},
			{
				Name:        "db_proxy_arn",
				Description: "The Amazon Resource Name (ARN) for the proxy",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBProxy.DBProxyArn")},
			{
				Name:        "created_date",
				Description: "The date and time when the proxy was first created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DBProxy.CreatedDate")},
			{
				Name:        "status",
				Description: "The current status of this proxy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBProxy.Status")},
			{
				Name:        "debug_logging",
				Description: "Whether the proxy includes detailed information about SQL statements in its logs.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBProxy.DebugLogging")},
			{
				Name:        "endpoint",
				Description: "The endpoint that you can use to connect to the DB proxy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBProxy.Endpoint")},
			{
				Name:        "engine_family",
				Description: "The kinds of databases that the proxy can connect to.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBProxy.EngineFamily")},
			{
				Name:        "idle_client_timeout",
				Description: "The number of seconds a connection to the proxy can have no activity before the proxy drops the client connection.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DBProxy.IdleClientTimeout")},
			{
				Name:        "require_tls",
				Description: "Indicates whether Transport Layer Security (TLS) encryption is required for connections to the proxy.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DBProxy.RequireTLS")},
			{
				Name:        "role_arn",
				Description: "The Amazon Resource Name (ARN) for the IAM role that the proxy uses to access Amazon Secrets Manager.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBProxy.RoleArn")},
			{
				Name:        "updated_date",
				Description: "The date and time when the proxy was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DBProxy.UpdatedDate")},
			{
				Name:        "vpc_id",
				Description: "Provides the VPC ID of the DB proxy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBProxy.VpcId")},
			{
				Name:        "auth",
				Description: "One or more data structures specifying the authorization mechanism to connect to the associated RDS DB instance or Aurora DB cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBProxy.Auth")},
			{
				Name:        "vpc_security_group_ids",
				Description: "Provides a list of VPC security groups that the proxy belongs to.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBProxy.VpcSecurityGroupIds")},
			{
				Name:        "vpc_subnet_ids",
				Description: "The EC2 subnet IDs for the proxy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBProxy.VpcSubnetIds")},
			// Standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DBProxy.DBProxyName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DBProxy.DBProxyArn").Transform(arnToAkas),
			},
		}),
	}
}
