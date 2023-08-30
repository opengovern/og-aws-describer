package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsGlueConnection(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_glue_connection",
		Description: "AWS Glue Connection",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"EntityNotFoundException"}),
			},
			Hydrate: kaytu.GetGlueConnection,
		},
		List: &plugin.ListConfig{
			KeyColumns: plugin.OptionalColumns([]string{"connection_type"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidInputException"}),
			},
			Hydrate: kaytu.ListGlueConnection,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the connection definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Connection.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the connection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
			{
				Name:        "connection_type",
				Description: "The type of the connection. Currently, SFTP is not supported.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Connection.ConnectionType")},
			{
				Name:        "creation_time",
				Description: "The time that this connection definition was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Connection.CreationTime")},
			{
				Name:        "description",
				Description: "The description of the connection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Connection.Description")},
			{
				Name:        "last_updated_by",
				Description: "The user, group, or role that last updated this connection definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Connection.LastUpdatedBy")},
			{
				Name:        "last_updated_time",
				Description: "The last time that this connection definition was updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Connection.LastUpdatedTime")},
			{
				Name:        "connection_properties",
				Description: "These key-value pairs define parameters for the connection.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Connection.ConnectionProperties")},
			{
				Name:        "match_criteria",
				Description: "A list of criteria that can be used in selecting this connection.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Connection.MatchCriteria")},
			{
				Name:        "physical_connection_requirements",
				Description: "A map of physical connection requirements, such as virtual private cloud (VPC) and SecurityGroup, that are needed to make this connection successfully.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Connection.PhysicalConnectionRequirements")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Connection.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
