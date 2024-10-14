package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCodeBuildProject(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_codebuild_project",
		Description: "AWS CodeBuild Project",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidInputException"}),
			},
			Hydrate: opengovernance.GetCodeBuildProject,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCodeBuildProject,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the build project.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Project.Name"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the build project.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Project.Arn"),
			},
			{
				Name:        "description",
				Description: "A description that makes the build project easy to identify.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Project.Description"),
			},
			{
				Name:        "concurrent_build_limit",
				Description: "The maximum number of concurrent builds that are allowed for this project.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Project.ConcurrentBuildLimit"),
			},
			{
				Name:        "created",
				Description: "When the build project was created, expressed in Unix time format.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Project.Created"),
			},
			{
				Name:        "last_modified",
				Description: "When the build project's settings were last modified, expressed in Unix time format.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Project.LastModified"),
			},
			{
				Name:        "encryption_key",
				Description: "The AWS Key Management Service (AWS KMS) customer master key (CMK) to be.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Project.EncryptionKey"),
			},
			{
				Name:        "queued_timeout_in_minutes",
				Description: "The number of minutes a build is allowed to be queued before it times out.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Project.QueuedTimeoutInMinutes"),
			},
			{
				Name:        "service_role",
				Description: "The ARN of the AWS Identity and Access Management (IAM) role that enables AWS CodeBuild to interact with dependent AWS services on behalf of the AWS account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Project.ServiceRole"),
			},
			{
				Name:        "source_version",
				Description: "A version of the build input to be built for this project.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Project.SourceVersion"),
			},
			{
				Name:        "timeout_in_minutes",
				Description: "How long, in minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait before timing out any related build that did not get marked as completed.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Project.TimeoutInMinutes"),
			},
			{
				Name:        "artifacts",
				Description: "Information about the build output artifacts for the build project.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Project.Artifacts"),
			},
			{
				Name:        "badge",
				Description: "Information about the build badge for the build project.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Project.Badge"),
			},
			{
				Name:        "build_batch_config",
				Description: "A ProjectBuildBatchConfig object that defines the batch build options for the project.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Project.BuildBatchConfig"),
			},
			{
				Name:        "cache",
				Description: "Information about the cache for the build project.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Project.Cache"),
			},
			{
				Name:        "environment",
				Description: "Information about the build environment for this build project.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Project.Environment"),
			},
			{
				Name:        "file_system_locations",
				Description: "An array of ProjectFileSystemLocation objects for a CodeBuild build project.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Project.FileSystemLocations"),
			},
			{
				Name:        "logs_config",
				Description: "Information about logs for the build project. A project can create logs in Amazon CloudWatch Logs, an S3 bucket or both.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Project.LogsConfig"),
			},
			{
				Name:        "project_visibility",
				Description: "Visibility of the build project.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Project.ProjectVisibility"),
			},
			{
				Name:        "secondary_artifacts",
				Description: "An array of ProjectArtifacts objects.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Project.SecondaryArtifacts"),
			},
			{
				Name:        "secondary_source_versions",
				Description: "An array of ProjectSource objects.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Project.SecondarySourceVersions"),
			},
			{
				Name:        "secondary_sources",
				Description: "An array of ProjectSource objects.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Project.SecondarySources"),
			},
			{
				Name:        "source",
				Description: "Information about the build input source code for this build project.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Project.Source"),
			},
			{
				Name:        "vpc_config",
				Description: "Information about the VPC configuration that AWS CodeBuild accesses.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Project.VpcConfig"),
			},
			{
				Name:        "webhook",
				Description: " Information about a webhook that connects repository events to a build project in AWS CodeBuild.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Project.Webhook"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tag key and value pairs associated with this build project.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Project.Tags"),
			},

			// Standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Project.Name"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description").Transform(codeBuildProjectTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Project.Arn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func codeBuildProjectTurbotTags(_ context.Context, d *transform.TransformData) (interface{},
	error) {
	data := d.HydrateItem.(opengovernance.CodeBuildProject).Description.Project

	if data.Tags == nil {
		return nil, nil
	}

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if data.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range data.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}

	}
	return turbotTagsMap, nil
}
