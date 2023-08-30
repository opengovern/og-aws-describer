package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEcrImage(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ecr_image",
		Description: "AWS ECR Image",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListECRImage,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "repository_name", Require: plugin.Optional},
				{Name: "registry_id", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "repository_name",
				Description: "The name of the repository.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.RepositoryName")},
			{
				Name:        "artifact_media_type",
				Description: "The artifact media type of the image.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.ArtifactMediaType")},
			{
				Name:        "image_digest",
				Description: "The sha256 digest of the image manifest.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.ImageDigest")},
			{
				Name:        "image_uri",
				Description: "The URI for the image.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ImageScanFinding.Uri")},
			{
				Name:        "image_manifest_media_type",
				Description: "The media type of the image manifest.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.ImageManifestMediaType")},
			{
				Name:        "image_pushed_at",
				Description: "The date and time, expressed in standard JavaScript date format, at which the current image was pushed to the repository.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Image.ImagePushedAt")},
			{
				Name:        "image_size_in_bytes",
				Description: "The size, in bytes, of the image in the repository.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Image.ImageSizeInBytes")},
			{
				Name:        "last_recorded_pull_time",
				Description: "The date and time, expressed in standard JavaScript date format, when Amazon ECR recorded the last image pull.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Image.LastRecordedPullTime")},
			{
				Name:        "registry_id",
				Description: "The Amazon Web Services account ID associated with the registry to which this image belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.RegistryId")},
			{
				Name:        "image_scan_findings_summary",
				Description: "A summary of the last completed image scan.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Image.ImageScanFindingsSummary")},
			{
				Name:        "image_scan_status",
				Description: "The current state of the scan.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Image.ImageScanStatus")},
			{
				Name:        "image_tags",
				Description: "The list of tags associated with this image.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Image.ImageTags")},
		}),
	}
}
