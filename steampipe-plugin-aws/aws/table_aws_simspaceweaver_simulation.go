package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSimSpaceWeaverSimulation(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_simspaceweaver_simulation",
		Description: "AWS SimSpace Weaver Simulation",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
			Hydrate: kaytu.GetSimSpaceWeaverSimulation,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListSimSpaceWeaverSimulation,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the simulation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Simulation.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the simulation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Simulation.Arn")},
			{
				Name:        "creation_time",
				Description: "The time when the simulation was created, expressed as the number of seconds and milliseconds in UTC since the Unix epoch (0:0:0.000, January 1, 1970).",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Simulation.CreationTime")},
			{
				Name:        "status",
				Description: "The current status of the simulation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Simulation.Status")},
			{
				Name:        "execution_id",
				Description: "A universally unique identifier (UUID) for this simulation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SimulationItem.ExecutionId")},
			{
				Name:        "maximum_duration",
				Description: "The maximum running time of the simulation, specified as a number of months (m or M), hours (h or H), or days (d or D).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SimulationItem.MaximumDuration")},
			{
				Name:        "role_arn",
				Description: "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role that the simulation assumes to perform actions.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SimulationItem.RoleArn")},
			{
				Name:        "schema_error",
				Description: "An error message that SimSpace Weaver returns only if there is a problem with the simulation schema.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SimulationItem.SchemaError")},
			{
				Name:        "live_simulation_state",
				Description: "A collection of additional state information, such as domain and clock configuration.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SimulationItem.LiveSimulationState")},
			{
				Name:        "logging_configuration",
				Description: "Settings that control how SimSpace Weaver handles your simulation log data.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SimulationItem.LoggingConfiguration")},
			{
				Name:        "schema_s3_location",
				Description: "The location of the simulation schema in Amazon Simple Storage Service (Amazon S3).",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SimulationItem.SchemaS3Location")},
			{
				Name:        "target_status",
				Description: "The desired status of the simulation.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Simulation.TargetStatus")},

			// Steampipe standard coulumns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Simulation.Name").Transform(arnToTitle),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Simulation.Arn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
