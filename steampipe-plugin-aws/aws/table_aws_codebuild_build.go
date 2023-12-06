package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsCodeBuildBuild(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_codebuild_build",
		Description: "AWS CodeBuild Build",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCodeBuildBuild,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "id",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The ARN of the build.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Build.Arn"),
			},
			{
				Name:        "id",
				Description: "The unique identifier of the  build.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Build.Id"),
			},
			{
				Name:        "build_batch_arn",
				Description: "The ARN of the batch build that this build is a member of, if applicable.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Build.BuildBatchArn"),
			},
			{
				Name:        "build_complete",
				Description: "Indicates if the build is complete.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Build.BuildComplete"),
			},
			{
				Name:        "build_number",
				Description: "The number of the build.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Build.BuildNumber"),
			},
			{
				Name:        "build_status",
				Description: "The status of the build.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Build.BuildStatus"),
			},
			{
				Name:        "current_phase",
				Description: "The current build phase.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Build.CurrentPhase"),
			},
			{
				Name:        "encryption_key",
				Description: "The Key Management Service customer master key (CMK) to be used for encrypting the build output artifacts.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Build.EncryptionKey"),
			},
			{
				Name:        "end_time",
				Description: "The date and time that the build process ended, expressed in Unix time format.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Build.EndTime"),
			},
			{
				Name:        "initiator",
				Description: "The entity that started the build.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Build.Initiator"),
			},
			{
				Name:        "project_name",
				Description: "The name of the build project.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Build.ProjectName"),
			},
			{
				Name:        "queued_timeout_in_minutes",
				Description: "Specifies the amount of time, in minutes, that a build is allowed to be queued before it times out.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Build.QueuedTimeoutInMinutes"),
			},
			{
				Name:        "source_version",
				Description: "The identifier of the version of the source code to be built.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Build.SourceVersion"),
			},
			{
				Name:        "start_time",
				Description: "The date and time that the build started.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Build.StartTime"),
			},
			{
				Name:        "timeout_in_minutes",
				Description: "How long, in minutes, for CodeBuild to wait before timing out this build if it does not get marked as completed.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Build.TimeoutInMinutes"),
			},
			{
				Name:        "resolved_source_version",
				Description: "The identifier of the resolved version of this build's source code.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Build.ResolvedSourceVersion"),
			},
			{
				Name:        "artifacts",
				Description: "A BuildArtifacts object the defines the build artifacts for this build.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Build.Artifacts"),
			},
			{
				Name:        "cache",
				Description: "Information about the cache for the build.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Build.Cache"),
			},
			{
				Name:        "debug_session",
				Description: "Contains information about the debug session for this build.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Build.DebugSession"),
			},
			{
				Name:        "environment",
				Description: "Information about the build environment for this build project.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Build.Environment"),
			},
			{
				Name:        "exported_environment_variables",
				Description: "A list of exported environment variables for this build.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Build.ExportedEnvironmentVariables"),
			},
			{
				Name:        "file_system_locations",
				Description: "An array of ProjectFileSystemLocation objects for a CodeBuild build project.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Build.FileSystemLocations"),
			},
			{
				Name:        "logs",
				Description: "Information about the build's logs in CloudWatch Logs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Build.Logs"),
			},
			{
				Name:        "network_interfaces",
				Description: "Describes a network interface.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Build.NetworkInterface"),
			},
			{
				Name:        "phases",
				Description: "Information about all previous build phases that are complete and information about any current build phase that is not yet complete.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Build.Phases"),
			},
			{
				Name:        "report_arns",
				Description: "An array of the ARNs associated with this build's reports.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Build.ReportArns"),
			},
			{
				Name:        "secondary_artifacts",
				Description: "An array of BuildArtifacts objects the define the build artifacts for this build.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Build.SecondaryArtifacts"),
			},
			{
				Name:        "secondary_source_versions",
				Description: "An array of ProjectSourceVersion objects.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Build.SecondarySourceVersions"),
			},
			{
				Name:        "secondary_sources",
				Description: "An array of ProjectSource objects that define the sources for the build.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Build.SecondarySources"),
			},
			{
				Name:        "source",
				Description: "Information about the build input source code for the build project.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Build.Source"),
			},
			{
				Name:        "vpc_config",
				Description: "Information about the VPC configuration that CodeBuild accesses.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Build.VpcConfig"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Build.Arn"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Build.Arn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
