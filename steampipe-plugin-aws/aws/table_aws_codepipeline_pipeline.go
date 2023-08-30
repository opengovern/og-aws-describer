package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsCodepipelinePipeline(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_codepipeline_pipeline",
		Description: "AWS Codepipeline Pipeline",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"PipelineNotFoundException"}),
			},
			Hydrate: kaytu.GetCodePipelinePipeline,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCodePipelinePipeline,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the pipeline.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pipeline.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the pipeline.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Metadata.PipelineArn"),
			},
			{
				Name:        "created_at",
				Description: "The date and time the pipeline was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Metadata.Created")},
			{
				Name:        "role_arn",
				Description: "The Amazon Resource Name (ARN) for AWS CodePipeline to use to either perform actions with no actionRoleArn, or to use to assume roles for actions with an actionRoleArn.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pipeline.RoleArn")},
			{
				Name:        "updated_at",
				Description: "The date and time of the last update to the pipeline.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Metadata.Updated")},
			{
				Name:        "version",
				Description: "The version number of the pipeline.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Pipeline.Version")},
			{
				Name:        "encryption_key",
				Description: "The encryption key used to encrypt the data in the artifact store, such as an AWS Key Management Service (AWS KMS) key. If this is undefined, the default key for Amazon S3 is used.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Pipeline.ArtifactStore.EncryptionKey")},
			{
				Name:        "artifact_stores",
				Description: "A mapping of artifactStore objects and their corresponding AWS Regions. There must be an artifact store for the pipeline Region and for each cross-region action in the pipeline.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Pipeline.ArtifactStores")},
			{
				Name:        "stages",
				Description: "The stage in which to perform the action.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Pipeline.Stages")},
			{
				Name:        "tags_src",
				Description: "A list of tag key and value pairs associated with this pipeline.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pipeline.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(codepipelineTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Metadata.PipelineArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func codepipelineTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.CodePipelinePipeline).Description.Tags

	if len(tags) <= 0 {
		return map[string]string{}, nil
	}

	// Mapping the resource tags inside turbotTags
	turbotTagsMap := map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return turbotTagsMap, nil
}
